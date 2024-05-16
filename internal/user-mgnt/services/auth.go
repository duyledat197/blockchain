package services

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/reddit/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	commonPb "openmyth/blockchain/idl/pb/common"
	pb "openmyth/blockchain/idl/pb/user"
	"openmyth/blockchain/internal/user-mgnt/entities"
	"openmyth/blockchain/internal/user-mgnt/repositories"
	"openmyth/blockchain/pkg/iface/pubsub"
	"openmyth/blockchain/pkg/xerror"
	"openmyth/blockchain/util"
)

type authService struct {
	userRepo   repositories.UserRepository
	publisher  pubsub.Publisher
	privateKey string

	pb.UnimplementedAuthServiceServer
}

// NewAuthService creates a new AuthService instance with the provided user repository.
//
// userRepo: the user repository for the AuthService.
// Returns an AuthServiceServer.
func NewAuthService(userRepo repositories.UserRepository, publisher pubsub.Publisher, privateKey string) pb.AuthServiceServer {
	return &authService{
		userRepo:   userRepo,
		publisher:  publisher,
		privateKey: privateKey,
	}
}

// Login handles the authentication of a user.
//
// ctx: the context for the authentication.
// req: the login request containing user credentials.
// Returns a LoginResponse and an error.
func (s *authService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.userRepo.FindUserByUsername(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, xerror.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "user or password is wrong")
		}
		return nil, status.Errorf(codes.Internal, "unable to find user: %v", err)
	}

	if err := util.CheckPassword(req.GetPassword(), user.HashedPassword); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "user or password is wrong")
	}

	tkn, err := util.GenerateToken(&jwt.StandardClaims{
		Id:      user.ID.Hex(),
		Subject: user.UserName,
		Issuer:  "openmyth",
	}, 24*time.Hour) // token expire after 1 day
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to generate token: %v", err)
	}

	return &pb.LoginResponse{
		Token: tkn,
		User:  userToPb(user),
	}, nil
}

// Register handles the registration of a new user.
//
// ctx: the context for the registration.
// req: the registration request containing user information.
// Returns a RegisterResponse and an error.
func (s *authService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	pwd, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to hash password: %v", err)
	}
	u, err := s.userRepo.FindUserByUsername(ctx, req.GetUsername())
	if err != nil && !errors.Is(err, xerror.ErrNotFound) {
		return nil, status.Errorf(codes.Internal, "unable to find user: %v", err)
	}

	if u != nil {
		return nil, status.Errorf(codes.AlreadyExists, "username already exists")
	}
	// generate new private key
	privateKey, privateKeyStr, err := util.NewPrivateKey()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to generate private key: %v", err)
	}

	publicKey, pubKey := util.PubKeyFromPrivKey(privateKey)
	walletAddress := crypto.PubkeyToAddress(*pubKey).Hex()
	user := &entities.User{
		UserName:       req.GetUsername(),
		HashedPassword: pwd,
		PrivateKey:     privateKeyStr,
		WalletAddress:  walletAddress,
		Nonce:          uuid.NewString(),
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create user: %v", err)
	}

	// send some native token to new wallet
	go func() {
		tx := &commonPb.Transaction{
			PrivKey: s.privateKey,
			To:      walletAddress,
			Amount:  "10",
		}

		b, err := proto.Marshal(tx)
		if err != nil {
			slog.Error("failed to marshal tx", slog.Any("err", err))
			return
		}
		if err := s.publisher.Publish(context.Background(), commonPb.TopicEvent_TOPIC_EVENT_SEND_MY_TOKEN_TRANSACTION.String(), &pubsub.Pack{
			Key: []byte(user.ID.Hex()),
			Msg: b,
		}); err != nil {
			slog.Error("failed to publish tx", slog.Any("err", err))
		}
	}()

	// send some my token to new wallet
	go func() {
		tx := &commonPb.Transaction{
			PrivKey: s.privateKey,
			To:      walletAddress,
			Amount:  "1",
		}

		b, err := proto.Marshal(tx)
		if err != nil {
			slog.Error("failed to marshal tx", slog.Any("err", err))
			return
		}
		if err := s.publisher.Publish(context.Background(), commonPb.TopicEvent_TOPIC_EVENT_SEND_NATIVE_TOKEN_TRANSACTION.String(), &pubsub.Pack{
			Key: []byte(user.ID.Hex()),
			Msg: b,
		}); err != nil {
			slog.Error("failed to publish tx", slog.Any("err", err))
		}
	}()

	return &pb.RegisterResponse{
		PrivateKey: privateKeyStr,
		PublicKey:  publicKey,
	}, nil
}

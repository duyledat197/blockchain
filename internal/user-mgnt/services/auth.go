package services

import (
	"context"
	"errors"
	"time"

	"github.com/reddit/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "openmyth/blockchain/idl/pb/user"
	"openmyth/blockchain/internal/user-mgnt/entities"
	"openmyth/blockchain/internal/user-mgnt/repositories"
	"openmyth/blockchain/pkg/xerror"
	"openmyth/blockchain/util"
)

type authService struct {
	userRepo repositories.UserRepository

	pb.UnimplementedAuthServiceServer
}

func NewAuthService(userRepo repositories.UserRepository) pb.AuthServiceServer {
	return &authService{
		userRepo: userRepo,
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
		return nil, status.Errorf(codes.NotFound, "user or password is wrong")
	}

	tkn, err := util.GenerateToken(&jwt.StandardClaims{
		Id:      user.ID,
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
	if !errors.Is(err, xerror.ErrNotFound) {
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

	if err := s.userRepo.Create(ctx, &entities.User{
		UserName:       req.GetUsername(),
		HashedPassword: pwd,
		PrivateKey:     privateKeyStr,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create user: %v", err)
	}

	publicKey, _ := util.PubKeyFromPrivKey(privateKey)

	return &pb.RegisterResponse{
		PrivateKey: privateKeyStr,
		PublicKey:  publicKey,
	}, nil
}

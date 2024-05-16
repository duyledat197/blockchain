package services

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	commonPb "openmyth/blockchain/idl/pb/common"
	pb "openmyth/blockchain/idl/pb/user"
	"openmyth/blockchain/internal/user-mgnt/entities"
	"openmyth/blockchain/internal/user-mgnt/repositories"
	"openmyth/blockchain/pkg/xerror"
)

type userService struct {
	userRepo repositories.UserRepository
	pb.UnimplementedUserServiceServer
}

func NewUserService(userRepo repositories.UserRepository) pb.UserServiceServer {
	return &userService{
		userRepo: userRepo,
	}
}

// GetUserByID retrieves a user by ID.
//
// ctx: the context of the request.
// req: the request containing the user ID.
// Returns the user and an error.
func (s *userService) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	u, err := s.userRepo.FindUser(ctx, req.GetUserId())
	if err != nil {
		if errors.Is(err, xerror.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "unable to find user: %v", err)
	}

	return &pb.GetUserByIDResponse{
		Data: userToPb(u),
	}, nil
}

// GetUserPrivateKeyByID retrieves the private key of a user by user ID.
//
// ctx: the context of the request.
// req: the request containing the user ID.
// Returns the user's private key and an error.
func (s *userService) GetUserPrivateKeyByID(ctx context.Context, req *pb.GetUserPrivateKeyByIDRequest) (*pb.GetUserPrivateKeyByIDResponse, error) {
	u, err := s.userRepo.FindUser(ctx, req.GetUserId())
	if err != nil {
		if errors.Is(err, xerror.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "unable to find user: %v", err)
	}

	return &pb.GetUserPrivateKeyByIDResponse{
		PrivateKey: u.PrivateKey,
		Nonce:      u.Nonce,
	}, nil
}

func (s *userService) CreateUser(_ context.Context, _ *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *userService) GetList(_ context.Context, _ *pb.GetListUserRequest) (*pb.GetListUserResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *userService) UpdateUser(_ context.Context, _ *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	panic("not implemented") // TODO: Implement
}

func userToPb(e *entities.User) *commonPb.User {
	return &commonPb.User{
		Username:      e.UserName,
		Id:            e.ID.Hex(),
		WalletAddress: e.WalletAddress,
		Nonce:         e.Nonce,
	}
}

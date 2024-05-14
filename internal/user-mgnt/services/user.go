package services

import (
	"context"

	commonPb "openmyth/blockchain/idl/pb/common"
	pb "openmyth/blockchain/idl/pb/user"
	"openmyth/blockchain/internal/user-mgnt/entities"
)

type userService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() pb.UserServiceServer {
	return &userService{}
}

func (s *userService) GetUserByID(_ context.Context, _ *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	panic("not implemented") // TODO: Implement
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
		Username: e.UserName,
		Id:       e.ID,
	}
}

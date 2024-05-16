package services

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "openmyth/blockchain/idl/pb/user"
	"openmyth/blockchain/internal/user-mgnt/entities"
	"openmyth/blockchain/mocks"
	"openmyth/blockchain/pkg/xerror"
	"openmyth/blockchain/util"
)

func Test_authService_Register(t *testing.T) {
	t.Parallel()
	publisher := mocks.NewPublisher(t)
	type fields struct {
		userRepo  *mocks.UserRepository
		publisher *mocks.Publisher
		wg        *sync.WaitGroup
	}
	type args struct {
		ctx context.Context
		req *pb.RegisterRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
		setup   func(fields fields)
	}{
		// TODO: Add test cases.
		{
			name: "happy case",
			fields: fields{
				userRepo:  mocks.NewUserRepository(t),
				publisher: publisher,
				wg:        &sync.WaitGroup{},
			},
			args: args{
				ctx: context.Background(),
				req: &pb.RegisterRequest{
					Username: "test",
					Password: "test",
				},
			},
			setup: func(fields fields) {
				fields.userRepo.On("FindUserByUsername", mock.Anything, "test").Return(nil, xerror.ErrNotFound)
				fields.userRepo.On("Create", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					user := args.Get(1).(*entities.User)
					user.ID = primitive.NewObjectID()
				}).Return(nil)
				publisher.On("Publish", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "error username already exists",
			fields: fields{
				userRepo:  mocks.NewUserRepository(t),
				publisher: mocks.NewPublisher(t),
				wg:        &sync.WaitGroup{},
			},
			args: args{
				ctx: context.Background(),
				req: &pb.RegisterRequest{
					Username: "test",
					Password: "test",
				},
			},
			wantErr: status.Errorf(codes.AlreadyExists, "username already exists"),
			setup: func(fields fields) {
				fields.userRepo.On("FindUserByUsername", mock.Anything, "test").Return(&entities.User{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &authService{
				userRepo:  tt.fields.userRepo,
				publisher: tt.fields.publisher,
			}
			tt.setup(tt.fields)
			_, err := s.Register(tt.args.ctx, tt.args.req)
			tt.fields.wg.Wait()
			if tt.wantErr != nil {
				require.NotNil(t, err)
				require.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}

}

func Test_authService_Login(t *testing.T) {
	t.Parallel()
	type fields struct {
		userRepo *mocks.UserRepository
		wg       *sync.WaitGroup
	}
	type args struct {
		ctx context.Context
		req *pb.LoginRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
		setup   func(fields fields)
	}{
		// TODO: Add test cases.
		{
			name: "happy case",
			fields: fields{
				userRepo: mocks.NewUserRepository(t),
				wg:       &sync.WaitGroup{},
			},
			args: args{
				ctx: context.Background(),
				req: &pb.LoginRequest{
					Username: "test",
					Password: "test",
				},
			},
			setup: func(fields fields) {
				pwd, _ := util.HashPassword("test")
				fields.userRepo.On("FindUserByUsername", mock.Anything, "test").Return(&entities.User{
					HashedPassword: pwd,
				}, nil)
			},
		},
		{
			name: "error user not found",
			fields: fields{
				userRepo: mocks.NewUserRepository(t),
				wg:       &sync.WaitGroup{},
			},
			args: args{
				ctx: context.Background(),
				req: &pb.LoginRequest{
					Username: "test",
					Password: "test",
				},
			},
			wantErr: status.Errorf(codes.NotFound, "user or password is wrong"),
			setup: func(fields fields) {
				fields.userRepo.On("FindUserByUsername", mock.Anything, "test").Return(nil, xerror.ErrNotFound)
			},
		},
		{
			name: "error wrong password",
			fields: fields{
				userRepo: mocks.NewUserRepository(t),
				wg:       &sync.WaitGroup{},
			},
			args: args{
				ctx: context.Background(),
				req: &pb.LoginRequest{
					Username: "test",
					Password: "test",
				},
			},
			wantErr: status.Errorf(codes.InvalidArgument, "user or password is wrong"),
			setup: func(fields fields) {
				fields.userRepo.On("FindUserByUsername", mock.Anything, "test").Return(&entities.User{
					HashedPassword: "another-hashed-password",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &authService{
				userRepo: tt.fields.userRepo,
			}
			tt.setup(tt.fields)
			_, err := s.Login(tt.args.ctx, tt.args.req)
			tt.fields.wg.Wait()
			if tt.wantErr != nil {
				require.NotNil(t, err)
				require.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

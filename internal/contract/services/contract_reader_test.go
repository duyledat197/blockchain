package services

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "openmyth/blockchain/idl/pb/contract"
	userPb "openmyth/blockchain/idl/pb/user"
	"openmyth/blockchain/mocks"
	mtdt "openmyth/blockchain/pkg/metadata"
	"openmyth/blockchain/pkg/xerror"
)

func TestContractReaderService_SendTransactionV2(t *testing.T) {
	t.Parallel()

	type fields struct {
		myTokenRepo *mocks.MyTokenRepository
		publisher   *mocks.Publisher
		userClient  *mocks.UserServiceClient
	}
	type args struct {
		ctx context.Context
		req *pb.SendTransactionV2Request
	}

	md := mtdt.ImportUserInfoToCtx(&mtdt.Payload{UserID: "user-id"})
	ctx := metadata.NewIncomingContext(context.Background(), md)
	var tests = []struct {
		name    string
		args    args
		fields  fields
		wantErr error

		setup func(fields fields)
	}{
		{
			name: "happy case",
			args: args{
				ctx: ctx,
				req: &pb.SendTransactionV2Request{
					Signature: "0xc83d417a3b99535e784a72af0d9772c019c776aa0dfe4313c001a5548f6cf254477f5334c30da59531bb521278edc98f1959009253dda4ee9f63fe5562ead5aa1b",
					To:        "0x0d3d7a5d1a8fa8e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e",
					Amount:    "1",
				},
			},
			fields: fields{
				myTokenRepo: mocks.NewMyTokenRepository(t),
				publisher:   mocks.NewPublisher(t),
				userClient:  mocks.NewUserServiceClient(t),
			},
			setup: func(fields fields) {

				fields.userClient.On("GetUserPrivateKeyByID", mock.Anything, mock.Anything).Return(&userPb.GetUserPrivateKeyByIDResponse{
					PrivateKey: "ae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce",
					Nonce:      "bou",
				}, nil)

				fields.myTokenRepo.On("BalanceOf", mock.Anything).Return(big.NewInt(100), nil)

				fields.publisher.On("Publish", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "error unauthenticated",
			args: args{
				ctx: context.Background(),
				req: &pb.SendTransactionV2Request{
					Signature: "0xc83d417a3b99535e784a72af0d9772c019c776aa0dfe4313c001a5548f6cf254477f5334c30da59531bb521278edc98f1959009253dda4ee9f63fe5562ead5aa1b",
					To:        "0x0d3d7a5d1a8fa8e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e",
					Amount:    "1",
				},
			},
			fields: fields{
				myTokenRepo: mocks.NewMyTokenRepository(t),
				publisher:   mocks.NewPublisher(t),
				userClient:  mocks.NewUserServiceClient(t),
			},
			wantErr: status.Errorf(codes.Unauthenticated, "unauthenticated"),
			setup: func(fields fields) {
			},
		},
		{
			name: "error amount is not valid",
			args: args{
				ctx: ctx,
				req: &pb.SendTransactionV2Request{
					Signature: "0xc83d417a3b99535e784a72af0d9772c019c776aa0dfe4313c001a5548f6cf254477f5334c30da59531bb521278edc98f1959009253dda4ee9f63fe5562ead5aa1b",
					To:        "0x0d3d7a5d1a8fa8e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e",
					Amount:    "0.1",
				},
			},
			fields: fields{
				myTokenRepo: mocks.NewMyTokenRepository(t),
				publisher:   mocks.NewPublisher(t),
				userClient:  mocks.NewUserServiceClient(t),
			},
			wantErr: status.Errorf(codes.InvalidArgument, "amount is not valid"),
			setup: func(fields fields) {
			},
		},
		{
			name: "error unable to get user by id",
			args: args{
				ctx: ctx,
				req: &pb.SendTransactionV2Request{
					Signature: "0xc83d417a3b99535e784a72af0d9772c019c776aa0dfe4313c001a5548f6cf254477f5334c30da59531bb521278edc98f1959009253dda4ee9f63fe5562ead5aa1b",
					To:        "0x0d3d7a5d1a8fa8e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e",
					Amount:    "1",
				},
			},
			fields: fields{
				myTokenRepo: mocks.NewMyTokenRepository(t),
				publisher:   mocks.NewPublisher(t),
				userClient:  mocks.NewUserServiceClient(t),
			},
			wantErr: status.Errorf(codes.Internal, "unable to get user: %v", xerror.ErrNotFound),
			setup: func(fields fields) {
				fields.userClient.On("GetUserPrivateKeyByID", mock.Anything, mock.Anything).Return(nil, xerror.ErrNotFound)
			},
		},
		{
			name: "error wrong signature",
			args: args{
				ctx: ctx,
				req: &pb.SendTransactionV2Request{
					Signature: "0xc83d417a3b99535e784a72af0d9772c019c776aa0dfe4313c001a5548f6cf254477f5334c30da59531bb521278edc98f1959009253dda4ee9f63fe5562ead5213123",
					To:        "0x0d3d7a5d1a8fa8e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e",
					Amount:    "1",
				},
			},
			fields: fields{
				myTokenRepo: mocks.NewMyTokenRepository(t),
				publisher:   mocks.NewPublisher(t),
				userClient:  mocks.NewUserServiceClient(t),
			},
			wantErr: status.Errorf(codes.InvalidArgument, "signature is not valid"),
			setup: func(fields fields) {
				fields.userClient.On("GetUserPrivateKeyByID", mock.Anything, mock.Anything).Return(&userPb.GetUserPrivateKeyByIDResponse{
					PrivateKey: "ae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce",
					Nonce:      "bou",
				}, nil)
			},
		},
		{
			name: "error failed to retrieve balance",
			args: args{
				ctx: ctx,
				req: &pb.SendTransactionV2Request{
					Signature: "0xc83d417a3b99535e784a72af0d9772c019c776aa0dfe4313c001a5548f6cf254477f5334c30da59531bb521278edc98f1959009253dda4ee9f63fe5562ead5aa1b",
					To:        "0x0d3d7a5d1a8fa8e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e",
					Amount:    "1",
				},
			},
			fields: fields{
				myTokenRepo: mocks.NewMyTokenRepository(t),
				publisher:   mocks.NewPublisher(t),
				userClient:  mocks.NewUserServiceClient(t),
			},
			wantErr: status.Errorf(codes.Internal, "unable to get balance: %v", fmt.Errorf("something error")),
			setup: func(fields fields) {

				fields.userClient.On("GetUserPrivateKeyByID", mock.Anything, mock.Anything).Return(&userPb.GetUserPrivateKeyByIDResponse{
					PrivateKey: "ae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce",
					Nonce:      "bou",
				}, nil)

				fields.myTokenRepo.On("BalanceOf", mock.Anything).Return(nil, fmt.Errorf("something error"))

			},
		},
		{
			name: "error balance not enough",
			args: args{
				ctx: ctx,
				req: &pb.SendTransactionV2Request{
					Signature: "0xc83d417a3b99535e784a72af0d9772c019c776aa0dfe4313c001a5548f6cf254477f5334c30da59531bb521278edc98f1959009253dda4ee9f63fe5562ead5aa1b",
					To:        "0x0d3d7a5d1a8fa8e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e",
					Amount:    "1000",
				},
			},
			fields: fields{
				myTokenRepo: mocks.NewMyTokenRepository(t),
				publisher:   mocks.NewPublisher(t),
				userClient:  mocks.NewUserServiceClient(t),
			},
			wantErr: status.Errorf(codes.InvalidArgument, "balance is not enough"),
			setup: func(fields fields) {

				fields.userClient.On("GetUserPrivateKeyByID", mock.Anything, mock.Anything).Return(&userPb.GetUserPrivateKeyByIDResponse{
					PrivateKey: "ae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce",
					Nonce:      "bou",
				}, nil)

				fields.myTokenRepo.On("BalanceOf", mock.Anything).Return(big.NewInt(100), nil)
			},
		},
		{
			name: "error publish failed",
			args: args{
				ctx: ctx,
				req: &pb.SendTransactionV2Request{
					Signature: "0xc83d417a3b99535e784a72af0d9772c019c776aa0dfe4313c001a5548f6cf254477f5334c30da59531bb521278edc98f1959009253dda4ee9f63fe5562ead5aa1b",
					To:        "0x0d3d7a5d1a8fa8e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e7b8a6d1b9c8c7a9d7e",
					Amount:    "1",
				},
			},
			fields: fields{
				myTokenRepo: mocks.NewMyTokenRepository(t),
				publisher:   mocks.NewPublisher(t),
				userClient:  mocks.NewUserServiceClient(t),
			},
			wantErr: status.Errorf(codes.Internal, "unable to send transaction: %v", fmt.Errorf("something error")),
			setup: func(fields fields) {

				fields.userClient.On("GetUserPrivateKeyByID", mock.Anything, mock.Anything).Return(&userPb.GetUserPrivateKeyByIDResponse{
					PrivateKey: "ae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce",
					Nonce:      "bou",
				}, nil)

				fields.myTokenRepo.On("BalanceOf", mock.Anything).Return(big.NewInt(100), nil)

				fields.publisher.On("Publish", mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("something error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ContractReaderService{
				myTokenRepo: tt.fields.myTokenRepo,
				publisher:   tt.fields.publisher,
				userClient:  tt.fields.userClient,
			}
			tt.setup(tt.fields)
			_, err := s.SendTransactionV2(tt.args.ctx, tt.args.req)
			if tt.wantErr != nil {
				require.NotNil(t, err)
				require.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

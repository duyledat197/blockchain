package services

import (
	"context"
	"reflect"
	"testing"

	pb "openmyth/blockchain/idl/pb/user"
	"openmyth/blockchain/internal/user-mgnt/repositories"
)

func Test_authService_Register(t *testing.T) {
	type fields struct {
		userRepo repositories.UserRepository
	}
	type args struct {
		ctx context.Context
		req *pb.RegisterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.RegisterResponse
		wantErr bool
		setup   func(fields fields)
	}{
		// TODO: Add test cases.
		{
			name:    "happy case",
			fields:  fields{},
			args:    args{},
			want:    &pb.RegisterResponse{},
			wantErr: false,
			setup: func(fields fields) {
			},
		},
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &authService{
				userRepo: tt.fields.userRepo,
			}
			tt.setup(tt.fields)
			got, err := s.Register(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("authService.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authService.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

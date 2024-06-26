// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"
	user "openmyth/blockchain/idl/pb/user"

	mock "github.com/stretchr/testify/mock"
)

// UserServiceServer is an autogenerated mock type for the UserServiceServer type
type UserServiceServer struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) CreateUser(_a0 context.Context, _a1 *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *user.CreateUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.CreateUserRequest) (*user.CreateUserResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.CreateUserRequest) *user.CreateUserResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.CreateUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.CreateUserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetList provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) GetList(_a0 context.Context, _a1 *user.GetListUserRequest) (*user.GetListUserResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetList")
	}

	var r0 *user.GetListUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetListUserRequest) (*user.GetListUserResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetListUserRequest) *user.GetListUserResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.GetListUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.GetListUserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) GetUserByID(_a0 context.Context, _a1 *user.GetUserByIDRequest) (*user.GetUserByIDResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *user.GetUserByIDResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserByIDRequest) (*user.GetUserByIDResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserByIDRequest) *user.GetUserByIDResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.GetUserByIDResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.GetUserByIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserPrivateKeyByID provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) GetUserPrivateKeyByID(_a0 context.Context, _a1 *user.GetUserPrivateKeyByIDRequest) (*user.GetUserPrivateKeyByIDResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetUserPrivateKeyByID")
	}

	var r0 *user.GetUserPrivateKeyByIDResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserPrivateKeyByIDRequest) (*user.GetUserPrivateKeyByIDResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserPrivateKeyByIDRequest) *user.GetUserPrivateKeyByIDResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.GetUserPrivateKeyByIDResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.GetUserPrivateKeyByIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) UpdateUser(_a0 context.Context, _a1 *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 *user.UpdateUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.UpdateUserRequest) (*user.UpdateUserResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.UpdateUserRequest) *user.UpdateUserResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.UpdateUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.UpdateUserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedUserServiceServer provides a mock function with given fields:
func (_m *UserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	_m.Called()
}

// NewUserServiceServer creates a new instance of UserServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserServiceServer {
	mock := &UserServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

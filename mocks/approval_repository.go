// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entities "openmyth/blockchain/internal/contract/entities"

	mock "github.com/stretchr/testify/mock"
)

// ApprovalRepository is an autogenerated mock type for the ApprovalRepository type
type ApprovalRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *ApprovalRepository) Create(_a0 context.Context, _a1 *entities.Approval) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Approval) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByOwner provides a mock function with given fields: ctx, owner
func (_m *ApprovalRepository) FindByOwner(ctx context.Context, owner string) ([]*entities.Approval, error) {
	ret := _m.Called(ctx, owner)

	if len(ret) == 0 {
		panic("no return value specified for FindByOwner")
	}

	var r0 []*entities.Approval
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*entities.Approval, error)); ok {
		return rf(ctx, owner)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*entities.Approval); ok {
		r0 = rf(ctx, owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Approval)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetList provides a mock function with given fields: _a0
func (_m *ApprovalRepository) GetList(_a0 context.Context) ([]*entities.Approval, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetList")
	}

	var r0 []*entities.Approval
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*entities.Approval, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*entities.Approval); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Approval)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewApprovalRepository creates a new instance of ApprovalRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewApprovalRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ApprovalRepository {
	mock := &ApprovalRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

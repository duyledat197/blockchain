// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	contract "openmyth/blockchain/idl/contracts"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// MyTokenRepository is an autogenerated mock type for the MyTokenRepository type
type MyTokenRepository struct {
	mock.Mock
}

// BalanceOf provides a mock function with given fields: addr
func (_m *MyTokenRepository) BalanceOf(addr common.Address) (*big.Int, error) {
	ret := _m.Called(addr)

	if len(ret) == 0 {
		panic("no return value specified for BalanceOf")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address) (*big.Int, error)); ok {
		return rf(addr)
	}
	if rf, ok := ret.Get(0).(func(common.Address) *big.Int); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterLogs provides a mock function with given fields: ctx, q
func (_m *MyTokenRepository) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	ret := _m.Called(ctx, q)

	if len(ret) == 0 {
		panic("no return value specified for FilterLogs")
	}

	var r0 []types.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery) ([]types.Log, error)); ok {
		return rf(ctx, q)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery) []types.Log); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.FilterQuery) error); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContractAddress provides a mock function with given fields:
func (_m *MyTokenRepository) GetContractAddress() common.Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetContractAddress")
	}

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// ParseApproval provides a mock function with given fields: log
func (_m *MyTokenRepository) ParseApproval(log types.Log) (*contract.MyTokenApproval, error) {
	ret := _m.Called(log)

	if len(ret) == 0 {
		panic("no return value specified for ParseApproval")
	}

	var r0 *contract.MyTokenApproval
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*contract.MyTokenApproval, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *contract.MyTokenApproval); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.MyTokenApproval)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseTransfer provides a mock function with given fields: log
func (_m *MyTokenRepository) ParseTransfer(log types.Log) (*contract.MyTokenTransfer, error) {
	ret := _m.Called(log)

	if len(ret) == 0 {
		panic("no return value specified for ParseTransfer")
	}

	var r0 *contract.MyTokenTransfer
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*contract.MyTokenTransfer, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *contract.MyTokenTransfer); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.MyTokenTransfer)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeFilterLogs provides a mock function with given fields: ctx, q, ch
func (_m *MyTokenRepository) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	ret := _m.Called(ctx, q, ch)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeFilterLogs")
	}

	var r0 ethereum.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) (ethereum.Subscription, error)); ok {
		return rf(ctx, q, ch)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) ethereum.Subscription); ok {
		r0 = rf(ctx, q, ch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ethereum.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) error); ok {
		r1 = rf(ctx, q, ch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transfer provides a mock function with given fields: ctx, privKey, toAdrr, amount
func (_m *MyTokenRepository) Transfer(ctx context.Context, privKey string, toAdrr string, amount *big.Int) error {
	ret := _m.Called(ctx, privKey, toAdrr, amount)

	if len(ret) == 0 {
		panic("no return value specified for Transfer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *big.Int) error); ok {
		r0 = rf(ctx, privKey, toAdrr, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMyTokenRepository creates a new instance of MyTokenRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMyTokenRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MyTokenRepository {
	mock := &MyTokenRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

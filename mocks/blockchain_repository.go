// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	ecdsa "crypto/ecdsa"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// BlockchainRepository is an autogenerated mock type for the BlockchainRepository type
type BlockchainRepository struct {
	mock.Mock
}

// RetrieveBalanceOf provides a mock function with given fields: ctx, address
func (_m *BlockchainRepository) RetrieveBalanceOf(ctx context.Context, address common.Address) (uint64, error) {
	ret := _m.Called(ctx, address)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveBalanceOf")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) (uint64, error)); ok {
		return rf(ctx, address)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) uint64); ok {
		r0 = rf(ctx, address)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveLatestBlock provides a mock function with given fields: ctx
func (_m *BlockchainRepository) RetrieveLatestBlock(ctx context.Context) (*types.Block, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveLatestBlock")
	}

	var r0 *types.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*types.Block, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *types.Block); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTransaction provides a mock function with given fields: ctx, privateKey, fromAddress, toAddress, value
func (_m *BlockchainRepository) SendTransaction(ctx context.Context, privateKey *ecdsa.PrivateKey, fromAddress common.Address, toAddress common.Address, value *big.Int) error {
	ret := _m.Called(ctx, privateKey, fromAddress, toAddress, value)

	if len(ret) == 0 {
		panic("no return value specified for SendTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecdsa.PrivateKey, common.Address, common.Address, *big.Int) error); ok {
		r0 = rf(ctx, privateKey, fromAddress, toAddress, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBlockchainRepository creates a new instance of BlockchainRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlockchainRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlockchainRepository {
	mock := &BlockchainRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

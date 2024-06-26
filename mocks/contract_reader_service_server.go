// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"
	contract "openmyth/blockchain/idl/pb/contract"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	mock "github.com/stretchr/testify/mock"
)

// ContractReaderServiceServer is an autogenerated mock type for the ContractReaderServiceServer type
type ContractReaderServiceServer struct {
	mock.Mock
}

// GetListApproval provides a mock function with given fields: _a0, _a1
func (_m *ContractReaderServiceServer) GetListApproval(_a0 context.Context, _a1 *emptypb.Empty) (*contract.GetListApprovalResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetListApproval")
	}

	var r0 *contract.GetListApprovalResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) (*contract.GetListApprovalResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) *contract.GetListApprovalResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.GetListApprovalResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListTransfer provides a mock function with given fields: _a0, _a1
func (_m *ContractReaderServiceServer) GetListTransfer(_a0 context.Context, _a1 *emptypb.Empty) (*contract.GetListTransferResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetListTransfer")
	}

	var r0 *contract.GetListTransferResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) (*contract.GetListTransferResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) *contract.GetListTransferResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.GetListTransferResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveBalanceOf provides a mock function with given fields: _a0, _a1
func (_m *ContractReaderServiceServer) RetrieveBalanceOf(_a0 context.Context, _a1 *contract.RetrieveBalanceOfRequest) (*contract.RetrieveBalanceOfResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveBalanceOf")
	}

	var r0 *contract.RetrieveBalanceOfResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *contract.RetrieveBalanceOfRequest) (*contract.RetrieveBalanceOfResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *contract.RetrieveBalanceOfRequest) *contract.RetrieveBalanceOfResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.RetrieveBalanceOfResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *contract.RetrieveBalanceOfRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveLatestBlock provides a mock function with given fields: _a0, _a1
func (_m *ContractReaderServiceServer) RetrieveLatestBlock(_a0 context.Context, _a1 *emptypb.Empty) (*contract.RetrieveLatestBlockResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveLatestBlock")
	}

	var r0 *contract.RetrieveLatestBlockResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) (*contract.RetrieveLatestBlockResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) *contract.RetrieveLatestBlockResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.RetrieveLatestBlockResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTransaction provides a mock function with given fields: _a0, _a1
func (_m *ContractReaderServiceServer) SendTransaction(_a0 context.Context, _a1 *contract.SendTransactionRequest) (*contract.SendTransactionResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SendTransaction")
	}

	var r0 *contract.SendTransactionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *contract.SendTransactionRequest) (*contract.SendTransactionResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *contract.SendTransactionRequest) *contract.SendTransactionResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.SendTransactionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *contract.SendTransactionRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTransactionV2 provides a mock function with given fields: _a0, _a1
func (_m *ContractReaderServiceServer) SendTransactionV2(_a0 context.Context, _a1 *contract.SendTransactionV2Request) (*contract.SendTransactionResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SendTransactionV2")
	}

	var r0 *contract.SendTransactionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *contract.SendTransactionV2Request) (*contract.SendTransactionResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *contract.SendTransactionV2Request) *contract.SendTransactionResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.SendTransactionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *contract.SendTransactionV2Request) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedContractReaderServiceServer provides a mock function with given fields:
func (_m *ContractReaderServiceServer) mustEmbedUnimplementedContractReaderServiceServer() {
	_m.Called()
}

// NewContractReaderServiceServer creates a new instance of ContractReaderServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewContractReaderServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *ContractReaderServiceServer {
	mock := &ContractReaderServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"
	contract "openmyth/blockchain/idl/pb/contract"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// ContractReaderServiceClient is an autogenerated mock type for the ContractReaderServiceClient type
type ContractReaderServiceClient struct {
	mock.Mock
}

// GetListApproval provides a mock function with given fields: ctx, in, opts
func (_m *ContractReaderServiceClient) GetListApproval(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*contract.GetListApprovalResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetListApproval")
	}

	var r0 *contract.GetListApprovalResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty, ...grpc.CallOption) (*contract.GetListApprovalResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty, ...grpc.CallOption) *contract.GetListApprovalResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.GetListApprovalResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListTransfer provides a mock function with given fields: ctx, in, opts
func (_m *ContractReaderServiceClient) GetListTransfer(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*contract.GetListTransferResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetListTransfer")
	}

	var r0 *contract.GetListTransferResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty, ...grpc.CallOption) (*contract.GetListTransferResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty, ...grpc.CallOption) *contract.GetListTransferResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.GetListTransferResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveBalanceOf provides a mock function with given fields: ctx, in, opts
func (_m *ContractReaderServiceClient) RetrieveBalanceOf(ctx context.Context, in *contract.RetrieveBalanceOfRequest, opts ...grpc.CallOption) (*contract.RetrieveBalanceOfResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveBalanceOf")
	}

	var r0 *contract.RetrieveBalanceOfResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *contract.RetrieveBalanceOfRequest, ...grpc.CallOption) (*contract.RetrieveBalanceOfResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *contract.RetrieveBalanceOfRequest, ...grpc.CallOption) *contract.RetrieveBalanceOfResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.RetrieveBalanceOfResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *contract.RetrieveBalanceOfRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveLatestBlock provides a mock function with given fields: ctx, in, opts
func (_m *ContractReaderServiceClient) RetrieveLatestBlock(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*contract.RetrieveLatestBlockResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveLatestBlock")
	}

	var r0 *contract.RetrieveLatestBlockResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty, ...grpc.CallOption) (*contract.RetrieveLatestBlockResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty, ...grpc.CallOption) *contract.RetrieveLatestBlockResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.RetrieveLatestBlockResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTransaction provides a mock function with given fields: ctx, in, opts
func (_m *ContractReaderServiceClient) SendTransaction(ctx context.Context, in *contract.SendTransactionRequest, opts ...grpc.CallOption) (*contract.SendTransactionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SendTransaction")
	}

	var r0 *contract.SendTransactionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *contract.SendTransactionRequest, ...grpc.CallOption) (*contract.SendTransactionResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *contract.SendTransactionRequest, ...grpc.CallOption) *contract.SendTransactionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.SendTransactionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *contract.SendTransactionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTransactionV2 provides a mock function with given fields: ctx, in, opts
func (_m *ContractReaderServiceClient) SendTransactionV2(ctx context.Context, in *contract.SendTransactionV2Request, opts ...grpc.CallOption) (*contract.SendTransactionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SendTransactionV2")
	}

	var r0 *contract.SendTransactionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *contract.SendTransactionV2Request, ...grpc.CallOption) (*contract.SendTransactionResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *contract.SendTransactionV2Request, ...grpc.CallOption) *contract.SendTransactionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.SendTransactionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *contract.SendTransactionV2Request, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewContractReaderServiceClient creates a new instance of ContractReaderServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewContractReaderServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *ContractReaderServiceClient {
	mock := &ContractReaderServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

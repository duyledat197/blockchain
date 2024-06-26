// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.14.0
// source: contract/contract_reader.proto

package contract

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ContractReaderService_GetListApproval_FullMethodName     = "/contract.ContractReaderService/GetListApproval"
	ContractReaderService_GetListTransfer_FullMethodName     = "/contract.ContractReaderService/GetListTransfer"
	ContractReaderService_RetrieveLatestBlock_FullMethodName = "/contract.ContractReaderService/RetrieveLatestBlock"
	ContractReaderService_RetrieveBalanceOf_FullMethodName   = "/contract.ContractReaderService/RetrieveBalanceOf"
	ContractReaderService_SendTransaction_FullMethodName     = "/contract.ContractReaderService/SendTransaction"
	ContractReaderService_SendTransactionV2_FullMethodName   = "/contract.ContractReaderService/SendTransactionV2"
)

// ContractReaderServiceClient is the client API for ContractReaderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContractReaderServiceClient interface {
	GetListApproval(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetListApprovalResponse, error)
	GetListTransfer(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetListTransferResponse, error)
	RetrieveLatestBlock(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RetrieveLatestBlockResponse, error)
	RetrieveBalanceOf(ctx context.Context, in *RetrieveBalanceOfRequest, opts ...grpc.CallOption) (*RetrieveBalanceOfResponse, error)
	SendTransaction(ctx context.Context, in *SendTransactionRequest, opts ...grpc.CallOption) (*SendTransactionResponse, error)
	SendTransactionV2(ctx context.Context, in *SendTransactionV2Request, opts ...grpc.CallOption) (*SendTransactionResponse, error)
}

type contractReaderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContractReaderServiceClient(cc grpc.ClientConnInterface) ContractReaderServiceClient {
	return &contractReaderServiceClient{cc}
}

func (c *contractReaderServiceClient) GetListApproval(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetListApprovalResponse, error) {
	out := new(GetListApprovalResponse)
	err := c.cc.Invoke(ctx, ContractReaderService_GetListApproval_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractReaderServiceClient) GetListTransfer(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetListTransferResponse, error) {
	out := new(GetListTransferResponse)
	err := c.cc.Invoke(ctx, ContractReaderService_GetListTransfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractReaderServiceClient) RetrieveLatestBlock(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RetrieveLatestBlockResponse, error) {
	out := new(RetrieveLatestBlockResponse)
	err := c.cc.Invoke(ctx, ContractReaderService_RetrieveLatestBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractReaderServiceClient) RetrieveBalanceOf(ctx context.Context, in *RetrieveBalanceOfRequest, opts ...grpc.CallOption) (*RetrieveBalanceOfResponse, error) {
	out := new(RetrieveBalanceOfResponse)
	err := c.cc.Invoke(ctx, ContractReaderService_RetrieveBalanceOf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractReaderServiceClient) SendTransaction(ctx context.Context, in *SendTransactionRequest, opts ...grpc.CallOption) (*SendTransactionResponse, error) {
	out := new(SendTransactionResponse)
	err := c.cc.Invoke(ctx, ContractReaderService_SendTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractReaderServiceClient) SendTransactionV2(ctx context.Context, in *SendTransactionV2Request, opts ...grpc.CallOption) (*SendTransactionResponse, error) {
	out := new(SendTransactionResponse)
	err := c.cc.Invoke(ctx, ContractReaderService_SendTransactionV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContractReaderServiceServer is the server API for ContractReaderService service.
// All implementations must embed UnimplementedContractReaderServiceServer
// for forward compatibility
type ContractReaderServiceServer interface {
	GetListApproval(context.Context, *emptypb.Empty) (*GetListApprovalResponse, error)
	GetListTransfer(context.Context, *emptypb.Empty) (*GetListTransferResponse, error)
	RetrieveLatestBlock(context.Context, *emptypb.Empty) (*RetrieveLatestBlockResponse, error)
	RetrieveBalanceOf(context.Context, *RetrieveBalanceOfRequest) (*RetrieveBalanceOfResponse, error)
	SendTransaction(context.Context, *SendTransactionRequest) (*SendTransactionResponse, error)
	SendTransactionV2(context.Context, *SendTransactionV2Request) (*SendTransactionResponse, error)
	mustEmbedUnimplementedContractReaderServiceServer()
}

// UnimplementedContractReaderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContractReaderServiceServer struct {
}

func (UnimplementedContractReaderServiceServer) GetListApproval(context.Context, *emptypb.Empty) (*GetListApprovalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListApproval not implemented")
}
func (UnimplementedContractReaderServiceServer) GetListTransfer(context.Context, *emptypb.Empty) (*GetListTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListTransfer not implemented")
}
func (UnimplementedContractReaderServiceServer) RetrieveLatestBlock(context.Context, *emptypb.Empty) (*RetrieveLatestBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveLatestBlock not implemented")
}
func (UnimplementedContractReaderServiceServer) RetrieveBalanceOf(context.Context, *RetrieveBalanceOfRequest) (*RetrieveBalanceOfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveBalanceOf not implemented")
}
func (UnimplementedContractReaderServiceServer) SendTransaction(context.Context, *SendTransactionRequest) (*SendTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTransaction not implemented")
}
func (UnimplementedContractReaderServiceServer) SendTransactionV2(context.Context, *SendTransactionV2Request) (*SendTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTransactionV2 not implemented")
}
func (UnimplementedContractReaderServiceServer) mustEmbedUnimplementedContractReaderServiceServer() {}

// UnsafeContractReaderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContractReaderServiceServer will
// result in compilation errors.
type UnsafeContractReaderServiceServer interface {
	mustEmbedUnimplementedContractReaderServiceServer()
}

func RegisterContractReaderServiceServer(s grpc.ServiceRegistrar, srv ContractReaderServiceServer) {
	s.RegisterService(&ContractReaderService_ServiceDesc, srv)
}

func _ContractReaderService_GetListApproval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractReaderServiceServer).GetListApproval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractReaderService_GetListApproval_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractReaderServiceServer).GetListApproval(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractReaderService_GetListTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractReaderServiceServer).GetListTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractReaderService_GetListTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractReaderServiceServer).GetListTransfer(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractReaderService_RetrieveLatestBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractReaderServiceServer).RetrieveLatestBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractReaderService_RetrieveLatestBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractReaderServiceServer).RetrieveLatestBlock(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractReaderService_RetrieveBalanceOf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveBalanceOfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractReaderServiceServer).RetrieveBalanceOf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractReaderService_RetrieveBalanceOf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractReaderServiceServer).RetrieveBalanceOf(ctx, req.(*RetrieveBalanceOfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractReaderService_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractReaderServiceServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractReaderService_SendTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractReaderServiceServer).SendTransaction(ctx, req.(*SendTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractReaderService_SendTransactionV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTransactionV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractReaderServiceServer).SendTransactionV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractReaderService_SendTransactionV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractReaderServiceServer).SendTransactionV2(ctx, req.(*SendTransactionV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ContractReaderService_ServiceDesc is the grpc.ServiceDesc for ContractReaderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContractReaderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contract.ContractReaderService",
	HandlerType: (*ContractReaderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetListApproval",
			Handler:    _ContractReaderService_GetListApproval_Handler,
		},
		{
			MethodName: "GetListTransfer",
			Handler:    _ContractReaderService_GetListTransfer_Handler,
		},
		{
			MethodName: "RetrieveLatestBlock",
			Handler:    _ContractReaderService_RetrieveLatestBlock_Handler,
		},
		{
			MethodName: "RetrieveBalanceOf",
			Handler:    _ContractReaderService_RetrieveBalanceOf_Handler,
		},
		{
			MethodName: "SendTransaction",
			Handler:    _ContractReaderService_SendTransaction_Handler,
		},
		{
			MethodName: "SendTransactionV2",
			Handler:    _ContractReaderService_SendTransactionV2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract/contract_reader.proto",
}

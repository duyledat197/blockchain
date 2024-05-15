package services

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	commonPb "openmyth/blockchain/idl/pb/common"
	pb "openmyth/blockchain/idl/pb/contract"
	"openmyth/blockchain/internal/contract/entities"
	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/pkg/xerror"
)

type ContractReaderService struct {
	approvalRepo   repositories.ApprovalRepository
	transferRepo   repositories.TransferRepository
	blockchainRepo repositories.BlockchainRepository

	pb.UnimplementedContractReaderServiceServer
}

func NewContractReaderService(
	approvalRepo repositories.ApprovalRepository,
	transferRepo repositories.TransferRepository,
	blockchainRepo repositories.BlockchainRepository,
) pb.ContractReaderServiceServer {
	return &ContractReaderService{
		approvalRepo:   approvalRepo,
		transferRepo:   transferRepo,
		blockchainRepo: blockchainRepo,
	}
}

func (s *ContractReaderService) GetListApproval(ctx context.Context, _ *emptypb.Empty) (*pb.GetListApprovalResponse, error) {
	approvals, err := s.approvalRepo.GetList(ctx)
	if err != nil {
		if errors.Is(err, xerror.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "not found")
		}

		return nil, status.Errorf(codes.Internal, "unable to get list approval: %v", err)
	}

	return &pb.GetListApprovalResponse{
		Data: approvalListToPbList(approvals),
	}, nil
}

func (s *ContractReaderService) GetListTransfer(ctx context.Context, _ *emptypb.Empty) (*pb.GetListTransferResponse, error) {
	transfers, err := s.transferRepo.GetList(ctx)
	if err != nil {
		if errors.Is(err, xerror.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "not found")
		}
		return nil, status.Errorf(codes.Internal, "unable to get list transfer: %v", err)
	}

	return &pb.GetListTransferResponse{
		Data: transferListToPbList(transfers),
	}, nil
}

func (s *ContractReaderService) RetrieveBalanceOf(ctx context.Context, req *pb.RetrieveBalanceOfRequest) (*pb.RetrieveBalanceOfResponse, error) {
	balance, err := s.blockchainRepo.RetrieveBalanceOf(ctx, req.GetAddress())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to retrieve balance: %v", err)
	}

	return &pb.RetrieveBalanceOfResponse{
		Balance: balance,
	}, nil
}
func (s *ContractReaderService) RetrieveLatestBlock(context.Context, *emptypb.Empty) (*pb.RetrieveLatestBlockResponse, error) {
	block, err := s.blockchainRepo.RetrieveLatestBlock(context.Background())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to retrieve latest block: %v", err)
	}
	return &pb.RetrieveLatestBlockResponse{
		Number:    block.NumberU64(),
		Nonce:     block.Nonce(),
		Hash:      block.Hash().Hex(),
		GasLimit:  block.GasLimit(),
		GasUsed:   block.GasUsed(),
		Timestamp: int64(block.Time()),
	}, nil
}

func approvalToPb(src *entities.Approval) *commonPb.Approval {
	return &commonPb.Approval{
		Owner:       src.Owner,
		Spender:     src.Spender,
		Value:       src.Value,
		BlockNumber: src.BlockNumber,
		Timestamp:   src.Timestamp,
	}
}

func transferToPb(src *entities.Transfer) *commonPb.Transfer {
	return &commonPb.Transfer{
		From:        src.From,
		To:          src.To,
		Value:       src.Value,
		BlockNumber: src.BlockNumber,
		Timestamp:   src.Timestamp,
	}
}

func approvalListToPbList(src []*entities.Approval) []*commonPb.Approval {
	dst := make([]*commonPb.Approval, 0, len(src))
	for _, v := range src {
		dst = append(dst, approvalToPb(v))
	}

	return dst
}

func transferListToPbList(src []*entities.Transfer) []*commonPb.Transfer {
	dst := make([]*commonPb.Transfer, 0, len(src))
	for _, v := range src {
		dst = append(dst, transferToPb(v))
	}

	return dst
}

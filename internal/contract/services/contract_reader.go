package services

import (
	"context"
	"errors"
	"log/slog"

	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	commonPb "openmyth/blockchain/idl/pb/common"
	pb "openmyth/blockchain/idl/pb/contract"
	"openmyth/blockchain/internal/contract/entities"
	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/pkg/iface/pubsub"
	"openmyth/blockchain/pkg/xerror"
)

// ContractReaderService is responsible for reading contract data from the blockchain and the database
// and publishing the results to a topic.
type ContractReaderService struct {
	// approvalRepo is the repository for approving contract proposals.
	approvalRepo repositories.ApprovalRepository
	// transferRepo is the repository for tracking token transfers.
	transferRepo repositories.TransferRepository
	// blockchainRepo is the repository for interacting with the blockchain.
	blockchainRepo repositories.BlockchainRepository
	// publisher is the publisher for publishing events.
	publisher pubsub.Publisher

	pb.UnimplementedContractReaderServiceServer
}

// NewContractReaderService initializes a new ContractReaderService with the provided approval repository, transfer repository, blockchain repository, and publisher.
//
// Parameters:
// - approvalRepo: The approval repository for the ContractReaderService.
// - transferRepo: The transfer repository for the ContractReaderService.
// - blockchainRepo: The blockchain repository for the ContractReaderService.
// - publisher: The publisher for the ContractReaderService.
// Return type: pb.ContractReaderServiceServer.
func NewContractReaderService(
	approvalRepo repositories.ApprovalRepository,
	transferRepo repositories.TransferRepository,
	blockchainRepo repositories.BlockchainRepository,
	publisher pubsub.Publisher,
) pb.ContractReaderServiceServer {
	return &ContractReaderService{
		approvalRepo:   approvalRepo,
		transferRepo:   transferRepo,
		blockchainRepo: blockchainRepo,
		publisher:      publisher,
	}
}

// GetListApproval retrieves the list of approvals.
//
// - ctx: The context for retrieving the list of approvals.
// - _ : The empty request parameter.
// Return type: *pb.GetListApprovalResponse and error.
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

// GetListTransfer retrieves the list of transfers.
// - ctx: The context for retrieving the list of transfers.
// - _ : The empty request parameter.
// Return type: *pb.GetListTransferResponse and error.
func (s *ContractReaderService) GetListTransfer(ctx context.Context, _ *emptypb.Empty) (*pb.GetListTransferResponse, error) {
	transfers, err := s.transferRepo.GetList(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to get list transfer: %v", err)
	}

	slog.Info("len", slog.Any("length", len(transfers)))

	return &pb.GetListTransferResponse{
		Data: transferListToPbList(transfers),
	}, nil
}

// RetrieveBalanceOf retrieves the balance of the specified address from the blockchain repository.
//
// Parameters:
// - ctx: The context for retrieving the balance.
// - req: The RetrieveBalanceOfRequest containing the address for which to retrieve the balance.
// Return type: RetrieveBalanceOfResponse and error.
func (s *ContractReaderService) RetrieveBalanceOf(ctx context.Context, req *pb.RetrieveBalanceOfRequest) (*pb.RetrieveBalanceOfResponse, error) {
	balance, err := s.blockchainRepo.RetrieveBalanceOf(ctx, req.GetAddress())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to retrieve balance: %v", err)
	}

	return &pb.RetrieveBalanceOfResponse{
		Balance: balance,
	}, nil
}

// RetrieveLatestBlock retrieves the latest block information.
//
// Parameters:
// - ctx: The context for retrieving the latest block.
// - _ : The empty request parameter.
// Return type: RetrieveLatestBlockResponse and error.
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

// SendTransaction sends a transaction based on the provided request.
//
// Parameters:
// - ctx: The context for the transaction.
// - req: The SendTransactionRequest containing transaction details.
// Return type: SendTransactionResponse and error.
func (s *ContractReaderService) SendTransaction(ctx context.Context, req *pb.SendTransactionRequest) (*pb.SendTransactionResponse, error) {
	b, err := proto.Marshal(&commonPb.Transaction{
		PrivKey: req.GetPrivKey(),
		To:      req.GetTo(),
		Amount:  req.Amount,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to marshal transaction: %v", err)
	}

	if err := s.publisher.Publish(ctx, commonPb.TopicEvent_TOPIC_EVENT_SEND_TRANSACTION.String(), &pubsub.Pack{
		Key: []byte(req.GetPrivKey()),
		Msg: b,
	}); err != nil {
		log.Error("unable to send transaction", slog.Any("err", err))
		return nil, status.Errorf(codes.Internal, "unable to send transaction: %v", err)
	}

	return &pb.SendTransactionResponse{}, nil
}

// approvalToPb converts an Approval entity to a commonPb.Approval struct.
//
// Parameters:
// - src: The Approval entity to convert.
// Return type: *commonPb.Approval.
func approvalToPb(src *entities.Approval) *commonPb.Approval {
	return &commonPb.Approval{
		Owner:       src.Owner,
		Spender:     src.Spender,
		Value:       src.Value,
		BlockNumber: src.BlockNumber,
		Timestamp:   src.Timestamp,
	}
}

// transferToPb converts an entities.Transfer to a commonPb.Transfer.
//
// Parameters:
// - src: The source entities.Transfer to be converted.
// Return type: *commonPb.Transfer.
func transferToPb(src *entities.Transfer) *commonPb.Transfer {
	return &commonPb.Transfer{
		From:        src.From,
		To:          src.To,
		Value:       src.Value,
		BlockNumber: src.BlockNumber,
		Timestamp:   src.Timestamp,
	}
}

// approvalListToPbList converts a list of Approval entities to a list of commonPb.Approval structs.
//
// - src: The list of Approval entities to convert.
// Return type: []*commonPb.Approval.
func approvalListToPbList(src []*entities.Approval) []*commonPb.Approval {
	dst := make([]*commonPb.Approval, 0, len(src))
	for _, v := range src {
		dst = append(dst, approvalToPb(v))
	}

	return dst
}

// transferListToPbList converts a list of entities.Transfer to a list of commonPb.Transfer.
//
// Parameters:
// - src: The source list of entities.Transfer to be converted.
// Return type: A list of *commonPb.Transfer.
func transferListToPbList(src []*entities.Transfer) []*commonPb.Transfer {
	dst := make([]*commonPb.Transfer, 0, len(src))
	for _, v := range src {
		dst = append(dst, transferToPb(v))
	}

	slog.Info("dst len", slog.Any("length", len(dst)))

	return dst
}

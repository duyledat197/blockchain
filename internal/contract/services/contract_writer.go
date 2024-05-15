package services

import (
	"context"
	"log/slog"
	"math/big"
	"time"

	"google.golang.org/protobuf/proto"

	pb "openmyth/blockchain/idl/pb/common"
	"openmyth/blockchain/internal/contract/entities"
	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/pkg/iface/pubsub"
)

// ContractWriterService is responsible for managing the logic for writing contracts.
//
// Fields:
// - approvalRepo: The approval repository used by the ContractWriterService.
// - transferRepo: The transfer repository used by the ContractWriterService.
// - myTokenRepo: The MyToken repository used by the ContractWriterService.
type ContractWriterService struct {
	approvalRepo repositories.ApprovalRepository
	transferRepo repositories.TransferRepository
	myTokenRepo  repositories.MyTokenRepository
}

// NewContractWriterService initializes a new ContractWriterService with the provided approval repository, transfer repository, and MyToken repository.
//
// Parameters:
// - approvalRepo: The approval repository for the ContractWriterService.
// - transferRepo: The transfer repository for the ContractWriterService.
// - myTokenRepo: The MyToken repository for the ContractWriterService.
// Return type: *ContractWriterService.
func NewContractWriterService(
	approvalRepo repositories.ApprovalRepository,
	transferRepo repositories.TransferRepository,
	myTokenRepo repositories.MyTokenRepository,
) *ContractWriterService {
	return &ContractWriterService{
		approvalRepo: approvalRepo,
		transferRepo: transferRepo,
		myTokenRepo:  myTokenRepo,
	}
}

// Subscribe handles subscribing to different topics and delegates the handling based on the topic type.
//
// - ctx: The context for the subscription.
// - topic: The topic string to determine the handling.
// - msg: The message containing the subscription details.
// - tt: The timestamp for the subscription event.
// Return type: None.
func (s *ContractWriterService) Subscribe(ctx context.Context, topic string, msg *pubsub.Pack, tt time.Time) {
	switch topic {
	case pb.TopicEvent_TOPIC_EVENT_APPROVAL.String():
		s.handleApprovalEvent(ctx, msg, tt)
	case pb.TopicEvent_TOPIC_EVENT_TRANSFER.String():
		s.handleTransferEvent(ctx, msg, tt)
	case pb.TopicEvent_TOPIC_EVENT_SEND_TRANSACTION.String():
		s.handleSendTransactionEvent(ctx, msg, tt)
	}
}

// handleApprovalEvent handles the approval event by unmarshaling the message and creating an approval record.
//
// Parameters:
// - ctx: The context for the approval event.
// - msg: The message containing the approval details.
// - tt: The timestamp for the approval event.
// Return type: None.
func (s *ContractWriterService) handleApprovalEvent(ctx context.Context, msg *pubsub.Pack, tt time.Time) {
	var approval pb.Approval
	if err := proto.Unmarshal(msg.Msg, &approval); err != nil {
		slog.Error("failed to unmarshal approval", slog.Any("err", err))
		return
	}
	if err := s.approvalRepo.Create(ctx, &entities.Approval{
		Owner:       approval.Owner,
		Spender:     approval.Spender,
		Value:       approval.Value,
		BlockNumber: approval.BlockNumber,
		Timestamp:   tt.Unix(),
	}); err != nil {
		slog.Error("failed to create approval", slog.Any("err", err))
	}
}

// handleTransferEvent handles the transfer event by unmarshaling the message and creating a transfer record.
//
// Parameters:
// - ctx: The context for the transfer event.
// - msg: The message containing the transfer details.
// - tt: The timestamp for the transfer event.
// Return type: None.
func (s *ContractWriterService) handleTransferEvent(ctx context.Context, msg *pubsub.Pack, tt time.Time) {
	var transfer pb.Transfer
	if err := proto.Unmarshal(msg.Msg, &transfer); err != nil {
		slog.Error("failed to unmarshal transfer", slog.Any("err", err))
		return
	}

	if err := s.transferRepo.Create(ctx, &entities.Transfer{
		From:        transfer.From,
		To:          transfer.To,
		Value:       transfer.Value,
		BlockNumber: transfer.BlockNumber,
		Timestamp:   tt.Unix(),
	}); err != nil {
		slog.Error("failed to create transfer", slog.Any("err", err))
	}
}

// handleSendTransactionEvent handles the sending of a transaction event.
//
// Parameters:
// - ctx: The context for the transaction.
// - msg: The message containing the transaction details.
// Return type: None.
func (s *ContractWriterService) handleSendTransactionEvent(ctx context.Context, msg *pubsub.Pack, _ time.Time) {
	var tx pb.Transaction
	if err := proto.Unmarshal(msg.Msg, &tx); err != nil {
		slog.Error("failed to unmarshal tx", slog.Any("err", err))
		return
	}

	amount := new(big.Int)
	if _, ok := amount.SetString(tx.Amount, 10); !ok {
		slog.Error("amount is not valid", slog.String("amount", tx.Amount))
		return
	}

	if err := s.myTokenRepo.Transfer(ctx, tx.PrivKey, tx.To, amount); err != nil {
		slog.Error("failed to dispatch tx", slog.Any("err", err))
	}
}

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

type ContractWriterService struct {
	approvalRepo repositories.ApprovalRepository
	transferRepo repositories.TransferRepository
	myTokenRepo  repositories.MyTokenRepository
}

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

	if err := s.myTokenRepo.Transfer(ctx, tx.PrivKey, tx.From, tx.To, amount); err != nil {
		slog.Error("failed to dispatch tx", slog.Any("err", err))
	}
}

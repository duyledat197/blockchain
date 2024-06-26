package services

import (
	"context"
	"log/slog"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/protobuf/proto"

	pb "openmyth/blockchain/idl/pb/common"
	"openmyth/blockchain/internal/contract/entities"
	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/pkg/iface/pubsub"
)

// ContractWriterService is responsible for managing the logic for writing contracts.
type ContractWriterService struct {
	approvalRepo   repositories.ApprovalRepository
	transferRepo   repositories.TransferRepository
	myTokenRepo    repositories.MyTokenRepository
	blockchainRepo repositories.BlockchainRepository
}

// NewContractWriterService initializes a new ContractWriterService with the provided approval repository, transfer repository, and MyToken repository.
func NewContractWriterService(
	approvalRepo repositories.ApprovalRepository,
	transferRepo repositories.TransferRepository,
	myTokenRepo repositories.MyTokenRepository,
	blockchainRepo repositories.BlockchainRepository,
) *ContractWriterService {
	return &ContractWriterService{
		approvalRepo:   approvalRepo,
		transferRepo:   transferRepo,
		myTokenRepo:    myTokenRepo,
		blockchainRepo: blockchainRepo,
	}
}

// Subscribe handles subscribing to different topics and delegates the handling based on the topic type.
func (s *ContractWriterService) Subscribe(ctx context.Context, topic string, msg *pubsub.Pack, tt time.Time) {
	switch topic {
	case pb.TopicEvent_TOPIC_EVENT_APPROVAL.String():
		s.handleApprovalEvent(ctx, msg, tt)
	case pb.TopicEvent_TOPIC_EVENT_TRANSFER.String():
		s.handleTransferEvent(ctx, msg, tt)
	case pb.TopicEvent_TOPIC_EVENT_SEND_MY_TOKEN_TRANSACTION.String():
		s.handleSendMyTokenTransactionEvent(ctx, msg, tt)
	case pb.TopicEvent_TOPIC_EVENT_SEND_NATIVE_TOKEN_TRANSACTION.String():
		s.handleSendNativeTokenTransactionEvent(ctx, msg, tt)
	}
}

// handleApprovalEvent handles the approval event by unmarshaling the message and creating an approval record.
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

// handleSendMyTokenTransactionEvent handles the sending of a transaction event.
func (s *ContractWriterService) handleSendMyTokenTransactionEvent(ctx context.Context, msg *pubsub.Pack, _ time.Time) {
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

// handleSendNativeTokenTransactionEvent handles the sending of a native token transaction event.
func (s *ContractWriterService) handleSendNativeTokenTransactionEvent(ctx context.Context, msg *pubsub.Pack, _ time.Time) {
	var tx pb.Transaction
	if err := proto.Unmarshal(msg.Msg, &tx); err != nil {
		slog.Error("failed to unmarshal tx", slog.Any("err", err))
		return
	}
	privateKey, err := crypto.HexToECDSA(tx.PrivKey)
	if err != nil {
		slog.Error("failed to convert private key", slog.Any("err", err))
		return
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	toAddress := common.HexToAddress(tx.To)
	amount := new(big.Int)
	if _, ok := amount.SetString(tx.Amount, 10); !ok {
		slog.Error("amount is not valid", slog.String("amount", tx.Amount))
		return
	}
	if err := s.blockchainRepo.SendTransaction(ctx, privateKey, fromAddress, toAddress, amount); err != nil {
		slog.Error("failed to dispatch tx", slog.Any("err", err))
	}

	slog.Info("send tx success")
}

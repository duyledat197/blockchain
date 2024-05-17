package services

import (
	"context"
	"errors"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	commonPb "openmyth/blockchain/idl/pb/common"
	pb "openmyth/blockchain/idl/pb/contract"
	userPb "openmyth/blockchain/idl/pb/user"
	"openmyth/blockchain/internal/contract/entities"
	"openmyth/blockchain/internal/contract/repositories"
	"openmyth/blockchain/pkg/iface/pubsub"
	"openmyth/blockchain/pkg/metadata"
	"openmyth/blockchain/pkg/xerror"
	"openmyth/blockchain/util/eth_util"
)

// ContractReaderService is responsible for reading contract data from the blockchain and the database
// and publishing the results to a topic.
type ContractReaderService struct {
	approvalRepo   repositories.ApprovalRepository
	transferRepo   repositories.TransferRepository
	blockchainRepo repositories.BlockchainRepository

	myTokenRepo repositories.MyTokenRepository

	userClient userPb.UserServiceClient

	publisher pubsub.Publisher

	pb.UnimplementedContractReaderServiceServer
}

// NewContractReaderService initializes a new ContractReaderService with the provided approval repository, transfer repository, blockchain repository, and publisher.
func NewContractReaderService(
	approvalRepo repositories.ApprovalRepository,
	transferRepo repositories.TransferRepository,
	blockchainRepo repositories.BlockchainRepository,
	myTokenRepo repositories.MyTokenRepository,
	userClient userPb.UserServiceClient,
	publisher pubsub.Publisher,
) pb.ContractReaderServiceServer {
	return &ContractReaderService{
		approvalRepo:   approvalRepo,
		transferRepo:   transferRepo,
		blockchainRepo: blockchainRepo,
		myTokenRepo:    myTokenRepo,
		publisher:      publisher,
		userClient:     userClient,
	}
}

// GetListApproval retrieves the list of approvals.
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
func (s *ContractReaderService) RetrieveBalanceOf(ctx context.Context, _ *pb.RetrieveBalanceOfRequest) (*pb.RetrieveBalanceOfResponse, error) {
	userCtx, ok := metadata.ExtractUserInfoFromCtx(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	u, err := s.userClient.GetUserByID(ctx, &userPb.GetUserByIDRequest{
		UserId: userCtx.UserID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to get user: %v", err)
	}

	walletAddress := u.GetData().GetWalletAddress()
	addr := common.HexToAddress(walletAddress)
	nativeBalance, err := s.blockchainRepo.RetrieveBalanceOf(ctx, addr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to retrieve balance: %v", err)
	}

	balance, err := s.myTokenRepo.BalanceOf(addr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to retrieve balance: %v", err)
	}

	return &pb.RetrieveBalanceOfResponse{
		NativeBalance: nativeBalance,
		Balance:       balance.Uint64(),
	}, nil
}

// RetrieveLatestBlock retrieves the latest block information.
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
func (s *ContractReaderService) SendTransaction(ctx context.Context, req *pb.SendTransactionRequest) (*pb.SendTransactionResponse, error) {
	b, err := proto.Marshal(&commonPb.Transaction{
		PrivKey: req.GetPrivKey(),
		To:      req.GetTo(),
		Amount:  req.Amount,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to marshal transaction: %v", err)
	}

	if err := s.publisher.Publish(ctx, commonPb.TopicEvent_TOPIC_EVENT_SEND_MY_TOKEN_TRANSACTION.String(), &pubsub.Pack{
		Key: []byte(req.GetPrivKey()),
		Msg: b,
	}); err != nil {
		log.Error("unable to send transaction", slog.Any("err", err))
		return nil, status.Errorf(codes.Internal, "unable to send transaction: %v", err)
	}

	return &pb.SendTransactionResponse{}, nil
}

// SendTransactionV2 sends a transaction based on the provided request.
func (s *ContractReaderService) SendTransactionV2(ctx context.Context, req *pb.SendTransactionV2Request) (*pb.SendTransactionResponse, error) {
	userCtx, ok := metadata.ExtractUserInfoFromCtx(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	if userCtx.UserID == "" {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	var amount big.Int
	if _, ok := amount.SetString(req.Amount, 10); !ok {
		return nil, status.Errorf(codes.InvalidArgument, "amount is not valid")
	}

	resp, err := s.userClient.GetUserPrivateKeyByID(ctx, &userPb.GetUserPrivateKeyByIDRequest{
		UserId: userCtx.UserID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to get user: %v", err)
	}

	if !eth_util.VerifySignature(resp.GetPrivateKey(), req.GetSignature(), resp.GetNonce()) {
		return nil, status.Errorf(codes.InvalidArgument, "signature is not valid")
	}

	privateKey, err := crypto.HexToECDSA(resp.GetPrivateKey())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to get private key: %v", err)
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	balance, err := s.myTokenRepo.BalanceOf(fromAddress)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to get balance: %v", err)
	}

	if balance.Cmp(&amount) < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "balance is not enough")
	}

	b, err := proto.Marshal(&commonPb.Transaction{
		PrivKey: resp.GetPrivateKey(),
		To:      req.GetTo(),
		Amount:  req.Amount,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to marshal transaction: %v", err)
	}

	if err := s.publisher.Publish(ctx, commonPb.TopicEvent_TOPIC_EVENT_SEND_MY_TOKEN_TRANSACTION.String(), &pubsub.Pack{
		Key: []byte(userCtx.UserID),
		Msg: b,
	}); err != nil {
		log.Error("unable to send transaction", slog.Any("err", err))
		return nil, status.Errorf(codes.Internal, "unable to send transaction: %v", err)
	}

	return &pb.SendTransactionResponse{}, nil
}

// approvalToPb converts an Approval entity to a commonPb.Approval struct.
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
func approvalListToPbList(src []*entities.Approval) []*commonPb.Approval {
	dst := make([]*commonPb.Approval, 0, len(src))
	for _, v := range src {
		dst = append(dst, approvalToPb(v))
	}

	return dst
}

// transferListToPbList converts a list of entities.Transfer to a list of commonPb.Transfer.
func transferListToPbList(src []*entities.Transfer) []*commonPb.Transfer {
	dst := make([]*commonPb.Transfer, 0, len(src))
	for _, v := range src {
		dst = append(dst, transferToPb(v))
	}

	return dst
}

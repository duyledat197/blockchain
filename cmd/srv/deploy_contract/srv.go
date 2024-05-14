package deploy_contract

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/lmittmann/tint"
	"gopkg.in/yaml.v3"

	"openmyth/blockchain/config"
	"openmyth/blockchain/pkg/contract"
	"openmyth/blockchain/pkg/eth"
)

type Server struct {
	cfg *config.Config
}

func (s *Server) loadConfig() {
	s.cfg = config.LoadConfig()
}

func (s *Server) loadLogger() {
	var slogHandler slog.Handler
	switch os.Getenv("ENV") {
	case "prod", "stg":
		f, err := os.OpenFile("./logs/deploy_contract.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("unable to open log file output: %v", err)
		}
		slogHandler = slog.NewJSONHandler(f, nil)
	default:
		slogHandler = tint.NewHandler(os.Stdout, nil)
	}

	logger := slog.New(slogHandler)

	slog.SetDefault(logger)
}

func (s *Server) Run(ctx context.Context) {
	s.loadConfig()
	s.loadLogger()

	client := eth.NewDialClient(s.cfg.ETHClient.Address())

	importedPrivateKey, err := crypto.HexToECDSA(s.cfg.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	ethClient := eth.NewEthClient(client)

	contractAddress, err := contract.DeployMyTokenContract(ctx, ethClient, importedPrivateKey)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	log.Println("Contract address:", contractAddress.Hex())

	b, err := yaml.Marshal(map[string]string{
		"contract_address": contractAddress.Hex(),
	})
	if err != nil {
		log.Fatalln(err)
	}

	// wite to config package for later use
	if err := os.WriteFile("./config/contract_address.yaml", b, os.ModePerm); err != nil {
		log.Fatalf("Failed to write contract address: %v", err)
	}
}

package config

import (
	"log"
	"os"
	"path"

	"github.com/spf13/viper"
)

// Config represents the overall configuration structure.
type Config struct {
	PostgresDB  *Database `mapstructure:"postgres_db,omitempty"`
	MongoDB     *Database `mapstructure:"mongo_db,omitempty"`
	Kafka       *Endpoint `mapstructure:"kafka,omitempty"`
	ETHClient   *Endpoint `mapstructure:"eth_client,omitempty"`
	WsETHClient *Endpoint `mapstructure:"ws_eth_client,omitempty"`

	UserService           *Endpoint `mapstructure:"user_service,omitempty"`
	GatewayService        *Endpoint `mapstructure:"gateway_service,omitempty"`
	ContractReaderService *Endpoint `mapstructure:"contract_reader_service,omitempty"`

	SymetricKey   string `mapstructure:"symetric_key,omitempty"`
	FileLogOutPut string `mapstructure:"file_log_out_put,omitempty"`

	ChainURL   string `mapstructure:"chain_url,omitempty"`
	PrivateKey string `mapstructure:"private_key,omitempty"`

	ContractAddress string `mapstructure:"contract_address,omitempty"`
}

// LoadConfig loads the configuration from the specified file path and environment.
func LoadConfig() *Config {

	// Initialize an instance of the private config structure.
	var cfg Config
	service := os.Getenv("SERVICE")
	viper.AddConfigPath(path.Join("./config", service))
	viper.AddConfigPath(path.Join("./config", "common"))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("unable to read config file: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to read config file: %w", err)
	}

	return &cfg
}

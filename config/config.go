package config

import (
	"log"

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
	Frontend              *Endpoint `mapstructure:"frontend,omitempty"`

	SymetricKey   string `mapstructure:"symetric_key,omitempty"`
	FileLogOutPut string `mapstructure:"file_log_out_put,omitempty"`

	ChainURL   string `mapstructure:"chain_url,omitempty"`
	PrivateKey string `mapstructure:"private_key,omitempty"`

	ContractAddress string `mapstructure:"contract_address,omitempty"`
}

// LoadConfig loads the configuration from the specified file path and environment.
func LoadConfig() *Config {
	viper.AutomaticEnv()
	// Initialize an instance of the private config structure.
	var cfg Config
	viper.AddConfigPath("/app/config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("unable to read config file: %v", err)
	}

	v := viper.New()
	v.AddConfigPath("/app/common")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("unable to read common config file: %v", err)
	}

	if err := viper.MergeConfigMap(v.AllSettings()); err != nil {
		log.Fatalf("unable to merge config file: %v", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to read config file: %v", err)
	}

	return &cfg
}

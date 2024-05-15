package config

import (
	"embed"
	"log"

	"github.com/spf13/viper"
)

//go:embed *.yaml
var cfgFiles embed.FS

// Config represents the overall configuration structure.
type Config struct {
	PostgresDB *Database `mapstructure:"postgres_db,omitempty"`
	MongoDB    *Database `mapstructure:"mongo_db,omitempty"`
	Kafka      *Endpoint `mapstructure:"kafka,omitempty"`
	ETHClient  *Endpoint `mapstructure:"eth_client,omitempty"`

	UserService           *Endpoint `mapstructure:"user_service,omitempty"`
	GatewayService        *Endpoint `mapstructure:"gateway_service,omitempty"`
	ContractReaderService *Endpoint `mapstructure:"contract_reader_service,omitempty"`
	ContractWriterService *Endpoint `mapstructure:"contract_writer_service,omitempty"`

	SymetricKey   string `mapstructure:"symetric_key,omitempty"`
	FileLogOutPut string `mapstructure:"file_log_out_put,omitempty"`

	ChainURL   string `mapstructure:"chain_url,omitempty"`
	PrivateKey string `mapstructure:"private_key,omitempty"`
}

// LoadConfig loads the configuration from the specified file path and environment.
func LoadConfig() *Config {

	// Initialize an instance of the private config structure.
	var cfg Config
	viper.SetConfigType("yaml")

	files, err := cfgFiles.ReadDir(".")
	if err != nil {
		log.Fatalf("unable to read dir: %v", err)
	}

	for _, f := range files {
		fInfo, err := f.Info()
		if err != nil {
			log.Fatalf("unable to get file info: %v", err)
		}

		fs, err := cfgFiles.Open(fInfo.Name())
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}

		if err := viper.MergeConfig(fs); err != nil {
			log.Fatalf("unable to read config file: %v", err)
		}
	}

	// Unmarshal the configuration into the private config structure.
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to unmarshal config file: %v", err)
	}

	log.Printf("Config: %+v", cfg)

	// Create and return the public Config structure based on the private config.
	return &cfg
}

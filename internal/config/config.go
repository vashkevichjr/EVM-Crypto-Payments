package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config contains all application configuration
type Config struct {
	// Database
	PostgresHost     string `mapstructure:"postgres_host"`
	PostgresPort     string `mapstructure:"postgres_port"`
	PostgresUser     string `mapstructure:"postgres_user"`
	PostgresPassword string `mapstructure:"postgres_password"`
	PostgresDB       string `mapstructure:"postgres_db"`
	PostgresSSLMode  string `mapstructure:"postgres_sslmode"`

	// Server
	ServerHost string `mapstructure:"server_host"`
	ServerPort string `mapstructure:"server_port"`

	// Blockchain RPC (will be used in Phase 4)
	PolygonRPCURL         string `mapstructure:"polygon_rpc_url"`
	EthereumSepoliaRPCURL string `mapstructure:"ethereum_sepolia_rpc_url"`

	// Encryption
	EncryptionKey string `mapstructure:"encryption_key"`

	// Logging
	LogLevel string `mapstructure:"log_level"`
}

// Load loads configuration from environment variables and .env file
func Load() (*Config, error) {
	v := viper.New()

	// Set default values
	setDefaults(v)

	// Configure environment variables
	// Viper will automatically read UPPER_CASE env vars and map them to lowercase keys
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// Try to read .env file (optional, won't fail if file doesn't exist)
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	// Ignore error if .env file doesn't exist
	if err := v.ReadInConfig(); err != nil {
		// It's ok if .env file doesn't exist, we'll use env vars and defaults
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	// Unmarshal configuration into struct
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Validate required fields
	if err := validate(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// setDefaults sets default values for configuration
func setDefaults(v *viper.Viper) {
	// Database defaults
	v.SetDefault("postgres_host", "localhost")
	v.SetDefault("postgres_port", "5432")
	v.SetDefault("postgres_user", "gateway")
	v.SetDefault("postgres_password", "gateway_password")
	v.SetDefault("postgres_db", "gateway_db")
	v.SetDefault("postgres_sslmode", "disable")

	// Server defaults
	v.SetDefault("server_host", "0.0.0.0")
	v.SetDefault("server_port", "8080")

	// Blockchain RPC defaults (empty strings)
	v.SetDefault("polygon_rpc_url", "")
	v.SetDefault("ethereum_sepolia_rpc_url", "")

	// Logging defaults
	v.SetDefault("log_level", "info")
}

// validate validates required configuration fields
func validate(cfg *Config) error {
	if cfg.EncryptionKey == "" {
		return fmt.Errorf("encryption_key is required (32 bytes hex string)")
	}

	if len(cfg.EncryptionKey) != 64 {
		return fmt.Errorf("encryption_key must be 64 hex characters (32 bytes)")
	}

	return nil
}

// DSN returns PostgreSQL connection string
func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresDB,
		c.PostgresSSLMode,
	)
}

package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config holds all application configuration
type Config struct {
	// Server settings
	ServerPort string
	ServerHost string

	// Blockchain settings
	ChainID              int64
	RPCEndpoint          string
	AIOracleAdapterAddr  string
	ResolutionModuleAddr string
	TokenAddr            string
	MarketFactoryAddr    string

	// AI settings
	OpenAIAPIKey string
	OpenAIModel  string

	// External API settings
	BSCScanAPIKey string

	// Signer settings
	SignerPrivateKey string // For local/testing
	UseKMS           bool
	KMSKeyID         string
	KMSRegion        string

	// Bond settings
	DefaultBondAmount string // in HORIZON tokens (e.g., "1000000000000000000000" = 1000 HORIZON)

	// Operational settings
	ProposalTimeout      time.Duration
	MaxConcurrentMarkets int
	LogLevel             string

	// Security
	AllowedOrigins []string
}

// LoadFromEnv loads configuration from environment variables
func LoadFromEnv() (*Config, error) {
	cfg := &Config{
		// Defaults
		ServerPort:           getEnv("SERVER_PORT", "8080"),
		ServerHost:           getEnv("SERVER_HOST", "0.0.0.0"),
		ChainID:              getEnvInt64("CHAIN_ID", 56), // BSC mainnet
		RPCEndpoint:          getEnv("RPC_ENDPOINT", ""),
		AIOracleAdapterAddr:  getEnv("AI_ORACLE_ADAPTER_ADDR", ""),
		ResolutionModuleAddr: getEnv("RESOLUTION_MODULE_ADDR", ""),
		TokenAddr:            getEnv("TOKEN_ADDR", ""),
		MarketFactoryAddr:    getEnv("MARKET_FACTORY_ADDR", ""),
		OpenAIAPIKey:         getEnv("OPENAI_API_KEY", ""),
		OpenAIModel:          getEnv("OPENAI_MODEL", "gpt-4-turbo-preview"),
		BSCScanAPIKey:        getEnv("BSCSCAN_API_KEY", ""),
		SignerPrivateKey:     getEnv("SIGNER_PRIVATE_KEY", ""),
		UseKMS:               getEnvBool("USE_KMS", false),
		KMSKeyID:             getEnv("KMS_KEY_ID", ""),
		KMSRegion:            getEnv("KMS_REGION", "us-east-1"),
		DefaultBondAmount:    getEnv("DEFAULT_BOND_AMOUNT", "1000000000000000000000"), // 1000 HORIZON
		ProposalTimeout:      getEnvDuration("PROPOSAL_TIMEOUT", 5*time.Minute),
		MaxConcurrentMarkets: getEnvInt("MAX_CONCURRENT_MARKETS", 10),
		LogLevel:             getEnv("LOG_LEVEL", "info"),
		AllowedOrigins:       []string{"*"}, // Configure based on deployment
	}

	// Validate required fields
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}

// Validate checks if all required configuration is present
func (c *Config) Validate() error {
	if c.RPCEndpoint == "" {
		return fmt.Errorf("RPC_ENDPOINT is required")
	}
	if c.AIOracleAdapterAddr == "" {
		return fmt.Errorf("AI_ORACLE_ADAPTER_ADDR is required")
	}
	if c.ResolutionModuleAddr == "" {
		return fmt.Errorf("RESOLUTION_MODULE_ADDR is required")
	}
	if c.TokenAddr == "" {
		return fmt.Errorf("TOKEN_ADDR is required")
	}
	if c.MarketFactoryAddr == "" {
		return fmt.Errorf("MARKET_FACTORY_ADDR is required")
	}
	if c.OpenAIAPIKey == "" {
		return fmt.Errorf("OPENAI_API_KEY is required")
	}

	// Validate signer configuration
	if !c.UseKMS && c.SignerPrivateKey == "" {
		return fmt.Errorf("either SIGNER_PRIVATE_KEY or USE_KMS=true must be set")
	}
	if c.UseKMS && c.KMSKeyID == "" {
		return fmt.Errorf("KMS_KEY_ID is required when USE_KMS=true")
	}

	return nil
}

// Helper functions for environment variable parsing

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}
	return defaultVal
}

func getEnvInt64(key string, defaultVal int64) int64 {
	if val := os.Getenv(key); val != "" {
		if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
			return intVal
		}
	}
	return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	if val := os.Getenv(key); val != "" {
		if boolVal, err := strconv.ParseBool(val); err == nil {
			return boolVal
		}
	}
	return defaultVal
}

func getEnvDuration(key string, defaultVal time.Duration) time.Duration {
	if val := os.Getenv(key); val != "" {
		if duration, err := time.ParseDuration(val); err == nil {
			return duration
		}
	}
	return defaultVal
}

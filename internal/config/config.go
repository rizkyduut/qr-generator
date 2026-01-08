package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Security SecurityConfig `mapstructure:"security"`
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port         string `mapstructure:"port"`
	MaxBodySize  int64  `mapstructure:"max_body_size"` // in bytes
	ReadTimeout  int    `mapstructure:"read_timeout"`  // in seconds
	WriteTimeout int    `mapstructure:"write_timeout"` // in seconds
}

// SecurityConfig holds security-related configuration
type SecurityConfig struct {
	APIKey                string `mapstructure:"api_key"`
	EnableSecurityHeaders bool   `mapstructure:"enable_security_headers"`
}

// LoadConfig reads configuration from file and environment variables
func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Set defaults
	viper.SetDefault("server.port", ":8080")
	viper.SetDefault("server.max_body_size", 1048576) // 1MB
	viper.SetDefault("server.read_timeout", 10)
	viper.SetDefault("server.write_timeout", 10)
	viper.SetDefault("security.api_key", "")
	viper.SetDefault("security.enable_security_headers", true)

	// Enable environment variables
	viper.SetEnvPrefix("QR")
	viper.AutomaticEnv()

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found, using defaults and environment variables")
		} else {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	} else {
		log.Printf("Using config file: %s", viper.ConfigFileUsed())
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}

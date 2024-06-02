// config/config.go
package config

import (
	"encoding/json"
	"os"
)

// Config struct represents the configuration
type Config struct {
	DatabaseConfig DatabaseConfig `json:"database_config"`
	SwaggerConfig  SwaggerConfig  `json:"swagger_config"`
}

type DatabaseConfig struct {
	ConnectionString string `json:"connection_string"`
	Port             int    `json:"port"`
	User             string `json:"user"`
	Password         string `json:"password"`
	Database         string `json:"database"`
}

type SwaggerConfig struct {
	SwaggerHost     string `json:"swagger_host"`
	SwaggerVersion  string `json:"swagger_version"`
	SwaggerBasePath string `json:"swagger_base_path"`
}

// LoadConfig loads the configuration from a JSON file
func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

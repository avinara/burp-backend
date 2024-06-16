// config/config.go
package config

import (
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Config struct represents the configuration
type Config struct {
	DatabaseConfig    DatabaseConfig `json:"database_config"`
	SwaggerConfig     SwaggerConfig  `json:"swagger_config"`
	GoogleLoginConfig oauth2.Config
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

	config.GoogleLoginConfig = oauth2.Config{
		RedirectURL:  "http://localhost:8080/google_callback",
		ClientID:     "CLIENT ID",
		ClientSecret: "<CLIENT SERCRET>",
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}

	return config, nil
}

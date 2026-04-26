package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/casapps/casimg/src/mode"
	"github.com/casapps/casimg/src/paths"
	"gopkg.in/yaml.v3"
)

// Config holds all application configuration
type Config struct {
	Mode    mode.Mode
	Paths   *paths.Paths
	Server  ServerConfig  `yaml:"server"`
	Cluster ClusterConfig `yaml:"cluster"`
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
	Debug   bool   `yaml:"debug"`
}

// ClusterConfig holds cluster configuration
type ClusterConfig struct {
	Enabled bool `yaml:"enabled"`
}

// Load loads configuration from file with defaults
func Load(configPath string) (*Config, error) {
	cfg := &Config{
		Mode:  mode.Detect(),
		Paths: paths.GetDefault(),
		Server: ServerConfig{
			Address: "0.0.0.0",
			Port:    64580,
			Debug:   false,
		},
		Cluster: ClusterConfig{
			Enabled: false,
		},
	}

	serverYML := filepath.Join(configPath, "server.yml")
	if _, err := os.Stat(serverYML); err == nil {
		data, err := os.ReadFile(serverYML)
		if err != nil {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}

		if err := yaml.Unmarshal(data, cfg); err != nil {
			return nil, fmt.Errorf("failed to parse config: %w", err)
		}
	}

	return cfg, nil
}

// Save saves configuration to file
func (c *Config) Save(configPath string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	serverYML := filepath.Join(configPath, "server.yml")
	if err := os.MkdirAll(configPath, 0755); err != nil {
		return fmt.Errorf("failed to create config dir: %w", err)
	}

	if err := os.WriteFile(serverYML, data, 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}

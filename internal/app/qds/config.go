package qds

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v2"
	"nba-task-main/internal/nats"
)

type Config struct {
	Nats nats.Config `yaml:"nats"`
	DSN  string      `yaml:"postgres"`
}

func NewConfig() (*Config, error) {
	fileName := os.Getenv("CONFIG_PATH")
	if fileName == "" {
		fileName = "config.yaml"
	}

	b, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w, you can specify path to the config file via CONFIG_PATH variable", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse envrionment: %w", err)
	}

	return &cfg, nil
}

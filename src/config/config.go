package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	LogLevel string `json:"log_level"`
	Handler  string `json:"handler"`
}

func LoadConfig(filename string) (*Config, error) {
	cfg := &Config{}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
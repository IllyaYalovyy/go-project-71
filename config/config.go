package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	SourceAddress  string `json:"sourceAddress"`
	SourcePort     int    `json:"sourcePort"`
	DestinationUrl string `json:"destinationUrl"`
}

func FromJson(path string) (*Config, error) {
	cfgBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)

	err = json.Unmarshal(cfgBytes, cfg)

	return cfg, err
}

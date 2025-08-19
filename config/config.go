package config

import (
	"encoding/json"
	"fmt"
	"log"
)

type SampleAuthConfig struct {
	DB      DBConfig      `json:"db" yaml:"db"`
	Server  ServerConfig  `json:"app" yaml:"app"`
	Redis   RedisConfig   `json:"redis" yaml:"redis"`
}

func (SampleAuthConfig) configSignature() {}

func (cfg SampleAuthConfig) Print() {
	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal struct to JSON: %v", err)
	}

	fmt.Printf("loaded config: %v", string(jsonData))
}

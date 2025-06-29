package main

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	DB DB `envPrefix:"DB_"`
}

type DB struct {
	Port     string `env:"PORT"`
	Host     string `env:"HOST"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Name     string `env:"NAME"`
	SSLMode  string `env:"SSL_MODE"`
}

func loadConfigFromEnv() (*Config, error) {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("parse cfg as %T from env: %w", cfg, err)
	}

	return &cfg, nil
}

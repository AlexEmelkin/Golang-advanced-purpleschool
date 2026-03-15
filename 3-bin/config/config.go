package config

import (
	"os"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")
	if key == "" {
		panic("НЕ передан Key")
	}
	return &Config{
		Key: key,
	}
}

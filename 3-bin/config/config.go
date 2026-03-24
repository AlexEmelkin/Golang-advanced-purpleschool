package config

import (
	"fmt"
	"os"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")
	if key == "" {
		fmt.Println("НЕ передан Key")
		key = "$2a$10$DijyuNp8Aq/Zit7XtipBUONtqbdHK9JKGscqqCBJiBz2jGYxijwQy"
	}
	return &Config{
		Key: key,
	}
}

package config

import "os"

type Config struct {
	Port string
	Mode string // gin mode: debug / release / test
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = "debug"
	}

	return &Config{
		Port: port,
		Mode: mode,
	}
}
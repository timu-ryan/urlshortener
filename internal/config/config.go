package config

import "flag"

type Config struct {
	ServerAddress string
	BaseURL       string
}

func NewConfig() Config {
	var cfg Config
	flag.StringVar(&cfg.ServerAddress, "a", ":8080", "HTTP server address")
	flag.StringVar(&cfg.BaseURL, "b", "http://localhost:8080", "Base URL for shortened links")
	flag.Parse()
	return cfg
}

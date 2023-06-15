package config

import (
	"time"
)

type AppConfig struct {
	Host         string
	Port         int
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	Debug        bool
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBDatabase   string
}

func NewDefaultAppConfig() *AppConfig {
	return &AppConfig{
		Host:         "",
		Port:         8080,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Debug:        false,
	}
}

func (cfg *AppConfig) ApplyOptions(options ...func(*AppConfig)) *AppConfig {
	for _, option := range options {
		option(cfg)
	}

	return cfg
}

func WithHost(host string) func(*AppConfig) {
	return func(s *AppConfig) {
		s.Host = host
	}
}

func WithPort(port int) func(*AppConfig) {
	return func(s *AppConfig) {
		s.Port = port
	}
}

func WithWriteTimeout(writeTimeout time.Duration) func(*AppConfig) {
	return func(s *AppConfig) {
		s.WriteTimeout = writeTimeout
	}
}

func WithReadTimeout(readTimeout time.Duration) func(*AppConfig) {
	return func(s *AppConfig) {
		s.ReadTimeout = readTimeout
	}
}

func WithDebug(debug bool) func(*AppConfig) {
	return func(s *AppConfig) {
		s.Debug = debug
	}
}

package config

import (
	"errors"
	"fmt"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewDefaultDBConfig() *DBConfig {
	return &DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	}
}

func (cfg *DBConfig) ApplyOptions(options ...func(*DBConfig)) *DBConfig {
	for _, option := range options {
		option(cfg)
	}

	if err := cfg.validate(); err != nil {
		panic(err)
	}

	return cfg
}

func (cfg *DBConfig) validate() error {
	if cfg.Host == "" {
		return errors.New("DBConfig.Host not set")
	}

	if cfg.Port == "" {
		return errors.New("DBConfig.Port not set")
	}

	if cfg.User == "" {
		return errors.New("DBConfig.User not set")
	}

	if cfg.Password == "" {
		return errors.New("DBConfig.Password not set")
	}

	if cfg.Database == "" {
		return errors.New("DBConfig.Database not set")
	}

	return nil
}

func WithDBHost(dbHost string) func(*DBConfig) {
	return func(s *DBConfig) {
		s.Host = dbHost
	}
}

func WithDBPort(dbPort string) func(*DBConfig) {
	return func(s *DBConfig) {
		s.Port = dbPort
	}
}

func WithDBUser(dbUser string) func(*DBConfig) {
	return func(s *DBConfig) {
		s.User = dbUser
	}
}

func WithDBPassword(dbPassword string) func(*DBConfig) {
	return func(s *DBConfig) {
		s.Password = dbPassword
	}
}

func WithDBDatabase(dbDatabase string) func(*DBConfig) {
	return func(s *DBConfig) {
		s.Database = dbDatabase
	}
}

func (config *DBConfig) GetConnectionURL() string {
	// "postgres://username:password@localhost:5432/database_name"
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Password, config.Host, config.Port, config.Database)
}

func (config *DBConfig) GetConnectionURLWithoutCredentials() string {
	// "postgres://localhost:5432/database_name"
	return fmt.Sprintf("postgres://%s:%s/%s", config.Host, config.Port, config.Database)
}

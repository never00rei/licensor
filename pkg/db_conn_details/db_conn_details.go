package dbconndetails

import (
	"errors"
	"fmt"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func GetDBConfigFromEnv() (*DBConfig, error) {
	Host := os.Getenv("DB_HOST")
	if Host == "" {
		return nil, errors.New("DB_HOST not set")
	}

	Port := os.Getenv("DB_PORT")
	if Port == "" {
		return nil, errors.New("DB_PORT not set")
	}

	User := os.Getenv("DB_USER")
	if User == "" {
		return nil, errors.New("DB_USER not set")
	}

	Password := os.Getenv("DB_PASSWORD")
	if Password == "" {
		return nil, errors.New("DB_PASSWORD not set")
	}

	Database := os.Getenv("DB_DATABASE")
	if Database == "" {
		return nil, errors.New("DB_DATABASE not set")
	}

	return GetDBConfig(Host, Port, User, Password, Database), nil
}

func GetDBConfig(Host, Port, User, Password, Database string) *DBConfig {
	return &DBConfig{
		Host:     Host,
		Port:     Port,
		User:     User,
		Password: Password,
		Database: Database,
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

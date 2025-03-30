package db

import (
	"gorm.io/gorm"
	"time"
)

type ConnectionConfig struct {
	Driver   string
	Host     string
	Username string
	Password string
	DBName   string
	Port     int
	SSLMode  string
}

type ConnectionPoolConfig struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

func NewRDBMS(config ConnectionConfig, poolConfig ConnectionPoolConfig) (*gorm.DB, error) {
	if config.Driver == "postgres" {
		return NewPostgresDB(config, poolConfig)
	} else {
		return NewMySqlDB(config, poolConfig)
	}
}

package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App        AppConfig
	Otel       OtelConfig
	DB         DbConfig
	GrpcServer GrpcServerConfig
}

type AppConfig struct {
	Name    string
	Version string
	Env     string
}

type GrpcServerConfig struct {
	Driver          string
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	GracefulTimeout time.Duration
}

type OtelConfig struct {
	Host string
	Port string
}

type DbConfig struct {
	Driver          string
	Host            string
	Port            int
	Name            string
	Username        string
	Password        string
	SSLMode         string
	MaxConnOpen     int
	MaxConnIdle     int
	MaxConnLifetime time.Duration
	MaxConnIdletime time.Duration
	AutoMigrate     bool
}

func LoadConfig() (Config, error) {
	vpr := viper.New()

	vpr.SetConfigFile(".env")
	vpr.AddConfigPath("./../")
	vpr.AutomaticEnv()

	_ = vpr.ReadInConfig()

	return Config{
		App: AppConfig{
			Name:    vpr.GetString("app_name"),
			Version: vpr.GetString("app_version"),
			Env:     vpr.GetString("app_env"),
		},
		Otel: OtelConfig{
			Host: vpr.GetString("otel_host"),
			Port: vpr.GetString("otel_port"),
		},
		DB: DbConfig{
			Driver:          vpr.GetString("db_driver"),
			Host:            vpr.GetString("db_host"),
			Port:            vpr.GetInt("db_port"),
			Name:            vpr.GetString("db_name"),
			Username:        vpr.GetString("db_username"),
			Password:        vpr.GetString("db_password"),
			SSLMode:         vpr.GetString("db_sslmode"),
			MaxConnOpen:     vpr.GetInt("db_pool_max_open_conn"),
			MaxConnIdle:     vpr.GetInt("db_pool_max_idle_conn"),
			MaxConnLifetime: vpr.GetDuration("db_pool_max_lifetime_conn"),
			MaxConnIdletime: vpr.GetDuration("db_pool_max_idletime_conn"),
			AutoMigrate:     vpr.GetBool("db_automigrate"),
		},
		GrpcServer: GrpcServerConfig{
			Driver:          vpr.GetString("grpc_server_driver"),
			Port:            vpr.GetString("grpc_server_port"),
			ReadTimeout:     vpr.GetDuration("grpc_server_read_timeout"),
			WriteTimeout:    vpr.GetDuration("grpc_server_write_timeout"),
			GracefulTimeout: vpr.GetDuration("grpc_server_graceful_timeout"),
		},
	}, nil
}

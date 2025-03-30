package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	configEnv "hexagonal-go-grpc/config"
	infraDb "hexagonal-go-grpc/internal/adapters/infrastructure/db"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/types"
	"hexagonal-go-grpc/internal/adapters/util"
	"log/slog"
	"os"
)

func main() {
	cmd := flag.String("cmd", "up", "command to migrate")
	value := flag.Int("value", 0, "value for execute command")

	ctx := context.Background()

	config, err := configEnv.LoadConfig()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to load config", slog.String(types.Error, err.Error()))
	}

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			return a
		},
	}).WithAttrs([]slog.Attr{
		slog.String("service", config.App.Name),
		slog.String("with-release", config.App.Version),
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	// Start Initial Infrastructure
	// Set up OpenTelemetry.
	otelShutdown, err := infraObsrv.SetupOTelSDK(ctx, infraObsrv.OtelConfig{
		Host:           config.Otel.Host,
		Port:           config.Otel.Port,
		ServiceName:    config.App.Name,
		ServiceVersion: config.App.Version,
		Environment:    config.App.Env,
	})
	if err != nil {
		slog.ErrorContext(ctx, "Failed to initialize otel sdk", slog.String(types.Error, err.Error()))
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	ctx, span := infraObsrv.Tracer().Start(ctx, "cmd:migrate:migrate.go:main")
	defer span.End()
	slog.InfoContext(ctx, "Start migration")
	gormdb, err := infraDb.NewRDBMS(infraDb.ConnectionConfig{
		Driver:   config.DB.Driver,
		Host:     config.DB.Host,
		Username: config.DB.Username,
		Password: config.DB.Password,
		DBName:   config.DB.Name,
		Port:     config.DB.Port,
		SSLMode:  config.DB.SSLMode,
	}, infraDb.ConnectionPoolConfig{
		MaxOpenConns:    config.DB.MaxConnOpen,
		MaxIdleConns:    config.DB.MaxConnIdle,
		ConnMaxLifetime: config.DB.MaxConnLifetime,
		ConnMaxIdleTime: config.DB.MaxConnIdletime,
	})
	err = util.DBMigrate(ctx, gormdb, config.DB, *cmd, value)
	if err != nil {
		if err == migrate.ErrNoChange {
			slog.InfoContext(ctx, fmt.Sprintf("no change when run migration with cmd : %s and value : %d", *cmd, value))
		} else {
			slog.ErrorContext(ctx, fmt.Sprintf("Failed to migrate with cmd : %s and value : %d", "up", 0), slog.String(types.Error, err.Error()))
		}
	}

	slog.InfoContext(ctx, "End migration")
}

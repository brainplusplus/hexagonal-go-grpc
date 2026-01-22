package main

import (
	"context"
	"errors"
	"fmt"
	configEnv "hexagonal-go-grpc/config"
	infraDb "hexagonal-go-grpc/internal/adapters/infrastructure/db"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	customerGrpcHandlerPkg "hexagonal-go-grpc/internal/adapters/primary/grpc/implementation/customer"
	"hexagonal-go-grpc/internal/adapters/primary/grpc/server/echo"
	customerDbRepositoryPkg "hexagonal-go-grpc/internal/adapters/secondary/repository/db/implementation/customer"
	"hexagonal-go-grpc/internal/adapters/secondary/repository/db/manager"
	customerServicePkg "hexagonal-go-grpc/internal/adapters/service/implementation/customer"
	"hexagonal-go-grpc/internal/adapters/types"
	"hexagonal-go-grpc/internal/adapters/util"
	"syscall"

	"github.com/golang-migrate/migrate/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"

	"log/slog"
	"os"
	"os/signal"
)

func main() {
	var ctx = context.Background()

	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	config, err := configEnv.LoadConfig()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to load config", slog.String(types.Error, err.Error()))
		return
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
	ctx, span := infraObsrv.Tracer().Start(ctx, "cmd:app:main.go:main")
	defer span.End()

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
	if config.DB.AutoMigrate {
		err = util.DBMigrate(ctx, gormdb, config.DB, "up", 0)
		if err != nil {
			if err == migrate.ErrNoChange {
				slog.InfoContext(ctx, "no change when run migration")
			} else {
				slog.ErrorContext(ctx, fmt.Sprintf("Failed to migrate with cmd : %s and value : %d", "up", 0), slog.String(types.Error, err.Error()))
			}
		}
	}

	transactionManager := manager.NewTransactionManager(gormdb)

	customerDbRepository := customerDbRepositoryPkg.New(gormdb)

	customerService := customerServicePkg.New(transactionManager, customerDbRepository)

	customerGrpcHandler := customerGrpcHandlerPkg.New(customerService)

	// Init web server.
	httpServer := echo.NewServer(echo.Config{
		Port:            config.GrpcServer.Port,
		ReadTimeout:     config.GrpcServer.ReadTimeout,
		WriteTimeout:    config.GrpcServer.WriteTimeout,
		GracefulTimeout: config.GrpcServer.GracefulTimeout,
	})
	r := httpServer.Router()
	r.Use(otelecho.Middleware(config.App.Name))

	// ROUTE
	r.Use(echo.Recoverer)
	echo.NewHandler(echo.Parameters{
		CustomerHandler: customerGrpcHandler,
	}).Register(r)

	// Run web server.
	httpServerChan := httpServer.Run()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case err := <-httpServerChan:
		if err != nil {
			slog.ErrorContext(ctx, "Error in server grpc/http", slog.String(types.Error, err.Error()))
			return
		}
	case <-sigChan:
	}
}

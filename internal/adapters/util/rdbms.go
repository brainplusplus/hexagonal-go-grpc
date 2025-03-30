package util

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // import file driver
	"gorm.io/gorm"
	"hexagonal-go-grpc/config"
	"hexagonal-go-grpc/internal/adapters/types"
	"log/slog"
	"strings"
)

func GetDBFromContext(ctx context.Context, db *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(types.TxKey).(*gorm.DB); ok {
		return tx
	}
	return db
}

func DBMigrate(ctx context.Context, gormdb *gorm.DB, config config.DbConfig, cmd string, value interface{}) error {
	db, err := gormdb.DB()
	if err != nil {
		panic(err)
	}

	var driver database.Driver
	if config.Driver == "postgres" {
		driver, err = postgres.WithInstance(db, &postgres.Config{})
	} else {
		driver, err = mysql.WithInstance(db, &mysql.Config{})
	}
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations/"+config.Driver, config.Name, driver)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
	}

	cmd = strings.ToLower(cmd)

	switch cmd {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "migrate":
		err = m.Migrate(uint(value.(int)))
	case "force":
		err = m.Force(value.(int))
	case "steps":
		err = m.Steps(value.(int))
	}

	if err != nil {
		if err == migrate.ErrNoChange {
			slog.InfoContext(ctx, err.Error())
		} else {
			slog.ErrorContext(ctx, err.Error())
		}
	}
	return err
}

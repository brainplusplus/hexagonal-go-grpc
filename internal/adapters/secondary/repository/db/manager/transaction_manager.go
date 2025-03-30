package manager

import (
	"context"
	"gorm.io/gorm"
	"hexagonal-go-grpc/internal/adapters/types"
	"hexagonal-go-grpc/internal/ports/secondary/db/manager"
)

type TransactionManager struct {
	DB *gorm.DB
}

func NewTransactionManager(db *gorm.DB) manager.ITransactionManager {
	return &TransactionManager{DB: db}
}

func (tm *TransactionManager) WithTransaction(
	fn func(ctx context.Context, args ...any) ([]any, int, error),
	ctx context.Context,
	args ...any,
) ([]any, int, error) {
	tx := tm.DB.Begin()
	if tx.Error != nil {
		return nil, 500, tx.Error
	}

	ctx = context.WithValue(ctx, types.TxKey, tx)

	results, statusCode, err := fn(ctx, args...)
	if err != nil {
		tx.Rollback()
		return nil, statusCode, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, 500, err
	}

	return results, statusCode, nil
}

func (tm *TransactionManager) GetDBFromContext(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(types.TxKey).(*gorm.DB); ok {
		return tx
	}
	return tm.DB
}

package manager

import (
	"context"
)

type ITransactionManager interface {
	WithTransaction(
		fn func(ctx context.Context, args ...any) ([]any, int, error),
		ctx context.Context,
		args ...any,
	) ([]any, int, error)
}

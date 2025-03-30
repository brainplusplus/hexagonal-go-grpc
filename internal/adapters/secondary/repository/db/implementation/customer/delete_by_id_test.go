package customer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"hexagonal-go-grpc/internal/adapters/secondary/repository/db/container"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	"net/http"
	"testing"
)

func TestCustomerDbRepository_DeleteById(t *testing.T) {
	ctx := context.Background()
	gormdb, err := container.GetGorm()
	assert.NoError(t, err)
	customerDbRepository := New(gormdb)

	assert.NoError(t, err)

	statusCode, err := customerDbRepository.DeleteById(ctx, entity.CustomerIdRequest{
		Id: 1,
	})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)
}

package customer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"hexagonal-go-grpc/internal/adapters/secondary/repository/db/container"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	"net/http"
	"testing"
)

func TestCustomerDbRepository_GetList(t *testing.T) {
	ctx := context.Background()
	gormdb, err := container.GetGorm()
	assert.NoError(t, err)

	customerDbRepository := New(gormdb)

	assert.NoError(t, err)

	customers, statusCode, err := customerDbRepository.GetList(ctx, entity.CustomerListRequest{})
	firstCustomer := customers[0]
	assert.NoError(t, err)
	assert.Equal(t, 2, len(customers))
	assert.Equal(t, 1, int(firstCustomer.ID))
	assert.Equal(t, "John Doe", firstCustomer.Name)
	assert.Equal(t, http.StatusOK, statusCode)
}

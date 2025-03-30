package customer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"hexagonal-go-grpc/internal/adapters/secondary/repository/db/container"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	"net/http"
	"testing"
)

func TestCustomerDbRepository_Create(t *testing.T) {
	ctx := context.Background()
	gormdb, err := container.GetGorm()
	assert.NoError(t, err)
	customerDbRepository := New(gormdb)

	assert.NoError(t, err)

	actualCustomer, statusCode, err := customerDbRepository.Create(ctx, entity.CustomerCreateRequest{
		Name:     "Bowo",
		Email:    "bowo@gmail.com",
		Phone:    "08521111111",
		Gender:   "Male",
		IsActive: false,
	})
	assert.NoError(t, err)
	assert.Equal(t, "Bowo", actualCustomer.Name)
	assert.Equal(t, "bowo@gmail.com", actualCustomer.Email)
	assert.Equal(t, "08521111111", actualCustomer.Phone)
	assert.Equal(t, http.StatusOK, statusCode)
}

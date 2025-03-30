package customer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"hexagonal-go-grpc/internal/adapters/secondary/repository/db/container"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	"net/http"
	"testing"
)

func TestCustomerDbRepository_GetByEmailOrPhone(t *testing.T) {
	ctx := context.Background()
	gormdb, err := container.GetGorm()
	assert.NoError(t, err)
	customerDbRepository := New(gormdb)

	assert.NoError(t, err)

	actualCustomer, statusCode, err := customerDbRepository.GetByEmailOrPhone(ctx, entity.CustomerEmailOrPhoneRequest{
		Email: "johndoe@example.com",
		Phone: "1234567890",
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, int(actualCustomer.ID))
	assert.Equal(t, "John Doe", actualCustomer.Name)
	assert.Equal(t, http.StatusOK, statusCode)
}

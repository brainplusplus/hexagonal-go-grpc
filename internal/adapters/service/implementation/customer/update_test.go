package customer

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"hexagonal-go-grpc/internal/adapters/service/entity"
	"hexagonal-go-grpc/internal/ports/secondary/db"
	"hexagonal-go-grpc/internal/ports/secondary/db/manager"
)

func TestCustomerServiceImpl_Update_Success(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	request := entity.CustomerUpdateRequest{
		ID:    123,
		Name:  "Updated Name",
		Email: "updated@example.com",
	}

	expected := entity.Customer{
		ID:    123,
		Name:  "Updated Name",
		Email: "updated@example.com",
	}

	mockCustomerAcc.EXPECT().Update(mock.Anything, request).Return(expected, http.StatusOK, nil)

	customerService := New(mockTrxMgr, mockCustomerAcc)

	actual, statusCode, err := customerService.Update(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Email, actual.Email)
}

func TestCustomerServiceImpl_Update_NotFound(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	request := entity.CustomerUpdateRequest{
		ID:    123,
		Name:  "Updated Name",
		Email: "updated@example.com",
	}

	mockCustomerAcc.EXPECT().Update(mock.Anything, request).Return(entity.Customer{}, http.StatusNotFound, errors.New("customer not found"))

	customerService := New(mockTrxMgr, mockCustomerAcc)

	actual, statusCode, err := customerService.Update(ctx, request)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)
	assert.Empty(t, actual.ID)
}

func TestCustomerServiceImpl_Update_DBError(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	request := entity.CustomerUpdateRequest{
		ID:    123,
		Name:  "Updated Name",
		Email: "updated@example.com",
	}

	expectedErr := errors.New("DB Error")

	mockCustomerAcc.EXPECT().Update(mock.Anything, request).Return(entity.Customer{}, http.StatusInternalServerError, expectedErr)

	customerService := New(mockTrxMgr, mockCustomerAcc)

	actual, statusCode, err := customerService.Update(ctx, request)
	assert.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Empty(t, actual.ID)
}

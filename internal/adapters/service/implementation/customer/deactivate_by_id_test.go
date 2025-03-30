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

func TestCustomerServiceImpl_DeactivateById_Success(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	request := entity.CustomerIdRequest{
		Id: 123,
	}

	expected := entity.Customer{
		ID:       123,
		IsActive: false,
	}

	mockCustomerAcc.EXPECT().DeactivateById(mock.Anything, request).Return(expected, http.StatusOK, nil)

	customerService := New(mockTrxMgr, mockCustomerAcc)

	actual, statusCode, err := customerService.DeactivateById(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, expected.ID, actual.ID)
	assert.False(t, actual.IsActive)
}

func TestCustomerServiceImpl_DeactivateById_NotFound(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	request := entity.CustomerIdRequest{
		Id: 123,
	}

	mockCustomerAcc.EXPECT().DeactivateById(mock.Anything, request).Return(entity.Customer{}, http.StatusNotFound, errors.New("customer not found"))

	customerService := New(mockTrxMgr, mockCustomerAcc)

	actual, statusCode, err := customerService.DeactivateById(ctx, request)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)
	assert.Empty(t, actual.ID)
}

func TestCustomerServiceImpl_DeactivateById_DBError(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	request := entity.CustomerIdRequest{
		Id: 123,
	}

	expectedErr := errors.New("DB Error")

	mockCustomerAcc.EXPECT().DeactivateById(mock.Anything, request).Return(entity.Customer{}, http.StatusInternalServerError, expectedErr)

	customerService := New(mockTrxMgr, mockCustomerAcc)

	actual, statusCode, err := customerService.DeactivateById(ctx, request)
	assert.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Empty(t, actual.ID)
}

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

func TestCustomerServiceImpl_DeleteById_Success(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	request := entity.CustomerIdRequest{
		Id: 123,
	}

	mockCustomerAcc.EXPECT().DeleteById(mock.Anything, request).Return(http.StatusOK, nil)

	customerService := New(mockTrxMgr, mockCustomerAcc)

	statusCode, err := customerService.DeleteById(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)
}

func TestCustomerServiceImpl_DeleteById_NotFound(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	request := entity.CustomerIdRequest{
		Id: 123,
	}

	mockCustomerAcc.EXPECT().DeleteById(mock.Anything, request).Return(http.StatusNotFound, errors.New("customer not found"))

	customerService := New(mockTrxMgr, mockCustomerAcc)

	statusCode, err := customerService.DeleteById(ctx, request)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)
}

func TestCustomerServiceImpl_DeleteById_DBError(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	request := entity.CustomerIdRequest{
		Id: 123,
	}

	expectedErr := errors.New("DB Error")
	mockCustomerAcc.EXPECT().DeleteById(mock.Anything, request).Return(http.StatusInternalServerError, expectedErr)

	customerService := New(mockTrxMgr, mockCustomerAcc)

	statusCode, err := customerService.DeleteById(ctx, request)
	assert.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
}

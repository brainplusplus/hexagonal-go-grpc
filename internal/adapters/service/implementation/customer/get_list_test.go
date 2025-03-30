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

func TestCustomerServiceImpl_GetList_Success(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	request := entity.CustomerListRequest{
		Ids:    nil,
		Search: "",
	}

	expected := []entity.Customer{
		{ID: 1, Name: "John Doe", Email: "john@example.com", IsActive: true},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com", IsActive: false},
	}

	mockCustomerAcc.EXPECT().GetList(mock.Anything, request).Return(expected, http.StatusOK, nil)

	customerService := New(mockTrxMgr, mockCustomerAcc)

	actual, statusCode, err := customerService.GetList(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, expected, actual)
}

func TestCustomerServiceImpl_GetList_NotFound(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	request := entity.CustomerListRequest{
		Ids:    nil,
		Search: "",
	}

	mockCustomerAcc.EXPECT().GetList(mock.Anything, request).Return([]entity.Customer{}, http.StatusNotFound, errors.New("customers not found"))

	customerService := New(mockTrxMgr, mockCustomerAcc)

	actual, statusCode, err := customerService.GetList(ctx, request)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)
	assert.Empty(t, actual)
}

func TestCustomerServiceImpl_GetList_DBError(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	request := entity.CustomerListRequest{
		Ids:    nil,
		Search: "",
	}

	expectedErr := errors.New("DB Error")

	mockCustomerAcc.EXPECT().GetList(mock.Anything, request).Return([]entity.Customer{}, http.StatusInternalServerError, expectedErr)

	customerService := New(mockTrxMgr, mockCustomerAcc)

	actual, statusCode, err := customerService.GetList(ctx, request)
	assert.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Empty(t, actual)
}

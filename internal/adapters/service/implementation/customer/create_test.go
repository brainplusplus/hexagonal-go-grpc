package customer

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	errors2 "hexagonal-go-grpc/internal/adapters/errors"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	"hexagonal-go-grpc/internal/ports/secondary/db"
	"hexagonal-go-grpc/internal/ports/secondary/db/manager"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerServiceImpl_CreateCustomer_Success(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	name := "John Doe"
	email := "johndoe@example.com"

	request := entity.CustomerCreateRequest{
		Name:  name,
		Email: email,
	}

	expected := entity.Customer{
		Name:  name,
		Email: email,
	}

	mockCustomerAcc.EXPECT().Create(mock.Anything, request).Return(expected, http.StatusOK, nil)

	customerService := New(mockTrxMgr, mockCustomerAcc)

	actual, statusCode, err := customerService.Create(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Email, actual.Email)
}

func TestCustomerServiceImpl_CreateCustomer_InvalidEmail(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	name := "John Doe"
	email := "johndoe@example.com"

	request := entity.CustomerCreateRequest{
		Name:  name,
		Email: email,
	}

	expected := entity.Customer{
		Name:  name,
		Email: email,
	}
	mockCustomerAcc.EXPECT().Create(mock.Anything, mock.Anything).Return(entity.Customer{}, http.StatusInternalServerError, errors2.ErrInternalDB)

	customerService := New(mockTrxMgr, mockCustomerAcc)

	actual, statusCode, err := customerService.Create(ctx, request)
	assert.Error(t, err)
	assert.NotEqual(t, http.StatusOK, statusCode)
	assert.NotEqual(t, expected.Name, actual.Name)
	assert.NotEqual(t, expected.Email, actual.Email)
}

func TestCustomerServiceImpl_CreateCustomer_DBError(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	name := "John Doe"
	email := "johndoe@example.com"

	request := entity.CustomerCreateRequest{
		Name:  name,
		Email: email,
	}

	expected := entity.Customer{
		Name:  name,
		Email: email,
	}

	expectedErr := errors.New("DB Error")

	mockCustomerAcc.EXPECT().Create(mock.Anything, mock.Anything).Return(entity.Customer{}, http.StatusInternalServerError, expectedErr)

	customerService := New(mockTrxMgr, mockCustomerAcc)

	actual, statusCode, err := customerService.Create(ctx, request)
	assert.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
	assert.NotEqual(t, http.StatusOK, statusCode)
	assert.NotEqual(t, expected.Name, actual.Name)
	assert.NotEqual(t, expected.Email, actual.Email)
}

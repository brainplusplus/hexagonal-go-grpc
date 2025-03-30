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

func TestCustomerServiceImpl_Register_Success(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	request := entity.CustomerRegisterRequest{
		Name:   "John Doe",
		Email:  "johndoe@example.com",
		Phone:  "1234567890",
		Gender: "Male",
	}

	expected := entity.Customer{
		ID:       1,
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Phone:    "1234567890",
		Gender:   "Male",
		IsActive: false,
	}
	// Mock WithTransaction to return the expected result
	// Simulasi transaksi agar fungsi di dalamnya benar-benar dieksekusi
	mockTrxMgr.EXPECT().
		WithTransaction(mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(fn func(context.Context, ...interface{}) ([]interface{}, int, error), ctx context.Context, args ...interface{}) ([]interface{}, int, error) {
			return fn(ctx, args...)
		})

	// Mock GetByEmailOrPhone to return not found
	mockCustomerAcc.EXPECT().
		GetByEmailOrPhone(mock.Anything, mock.MatchedBy(func(req entity.CustomerEmailOrPhoneRequest) bool {
			return req.Email == request.Email && req.Phone == request.Phone
		})).
		Return(entity.Customer{}, http.StatusNotFound, errors.New("not found"))

	// Mock Create to return expected customer
	mockCustomerAcc.EXPECT().
		Create(mock.Anything, mock.MatchedBy(func(req entity.CustomerCreateRequest) bool {
			return req.Name == request.Name && req.Email == request.Email && req.Phone == request.Phone
		})).
		Return(expected, http.StatusOK, nil)

	customerService := New(mockTrxMgr, mockCustomerAcc)
	actual, statusCode, err := customerService.Register(ctx, request)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, expected.Name, actual.Name)
}
func TestCustomerServiceImpl_Register_AlreadyExists(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	request := entity.CustomerRegisterRequest{
		Name:   "John Doe",
		Email:  "johndoe@example.com",
		Phone:  "1234567890",
		Gender: "Male",
	}

	expected := entity.Customer{
		ID:       1,
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Phone:    "1234567890",
		Gender:   "Male",
		IsActive: false,
	}

	// Simulasi transaksi yang mengembalikan error karena customer sudah ada
	mockTrxMgr.EXPECT().
		WithTransaction(mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(fn func(context.Context, ...interface{}) ([]interface{}, int, error), ctx context.Context, args ...interface{}) ([]interface{}, int, error) {
			return fn(ctx, args...)
		})

	// Ekspektasi pemanggilan GetByEmailOrPhone yang mengembalikan status OK (customer sudah ada)
	mockCustomerAcc.EXPECT().
		GetByEmailOrPhone(mock.Anything, mock.MatchedBy(func(req entity.CustomerEmailOrPhoneRequest) bool {
			return req.Email == request.Email && req.Phone == request.Phone
		})).
		Return(expected, http.StatusOK, nil)

	customerService := New(mockTrxMgr, mockCustomerAcc)
	actual, statusCode, err := customerService.Register(ctx, request)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, statusCode)
	assert.Empty(t, actual)
}

func TestCustomerServiceImpl_Register_DBError(t *testing.T) {
	mockTrxMgr := manager.NewMockITransactionManager(t)
	mockCustomerAcc := db.NewMockICustomerDbRepository(t)
	ctx := context.Background()

	request := entity.CustomerRegisterRequest{
		Name:   "John Doe",
		Email:  "johndoe@example.com",
		Phone:  "1234567890",
		Gender: "Male",
	}

	expectedErr := errors.New("DB Error")

	mockTrxMgr.EXPECT().
		WithTransaction(mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(fn func(context.Context, ...interface{}) ([]interface{}, int, error), ctx context.Context, args ...interface{}) ([]interface{}, int, error) {
			return fn(ctx, args...)
		})

	mockCustomerAcc.EXPECT().GetByEmailOrPhone(mock.Anything, mock.Anything).Return(entity.Customer{}, http.StatusNotFound, errors.New("not found"))
	mockCustomerAcc.EXPECT().Create(mock.Anything, mock.Anything).Return(entity.Customer{}, http.StatusInternalServerError, expectedErr)

	customerService := New(mockTrxMgr, mockCustomerAcc)
	actual, statusCode, err := customerService.Register(ctx, request)

	assert.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Empty(t, actual)
}

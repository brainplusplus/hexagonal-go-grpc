package db

import (
	"context"
	"hexagonal-go-grpc/internal/adapters/service/entity"
)

type ICustomerDbRepository interface {
	GetList(ctx context.Context, request entity.CustomerListRequest) ([]entity.Customer, int, error)
	GetByEmailOrPhone(ctx context.Context, request entity.CustomerEmailOrPhoneRequest) (entity.Customer, int, error)
	GetById(ctx context.Context, request entity.CustomerIdRequest) (entity.Customer, int, error)
	Create(ctx context.Context, request entity.CustomerCreateRequest) (entity.Customer, int, error)
	Update(ctx context.Context, request entity.CustomerUpdateRequest) (entity.Customer, int, error)
	DeleteById(ctx context.Context, request entity.CustomerIdRequest) (int, error)
	ActivateById(ctx context.Context, request entity.CustomerIdRequest) (entity.Customer, int, error)
	DeactivateById(ctx context.Context, request entity.CustomerIdRequest) (entity.Customer, int, error)
}

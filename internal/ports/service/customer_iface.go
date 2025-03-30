package service

import (
	"context"
	"hexagonal-go-grpc/internal/adapters/service/entity"
)

type ICustomerService interface {
	GetList(ctx context.Context, request entity.CustomerListRequest) ([]entity.Customer, int, error)
	GetById(ctx context.Context, request entity.CustomerIdRequest) (entity.Customer, int, error)
	Create(ctx context.Context, request entity.CustomerCreateRequest) (entity.Customer, int, error)
	Update(ctx context.Context, request entity.CustomerUpdateRequest) (entity.Customer, int, error)
	DeleteById(ctx context.Context, request entity.CustomerIdRequest) (int, error)
	ActivateById(ctx context.Context, request entity.CustomerIdRequest) (entity.Customer, int, error)
	DeactivateById(ctx context.Context, request entity.CustomerIdRequest) (entity.Customer, int, error)
	Register(ctx context.Context, request entity.CustomerRegisterRequest) (entity.Customer, int, error)
}

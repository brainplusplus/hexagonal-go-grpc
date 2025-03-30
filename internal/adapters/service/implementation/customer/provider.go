package customer

import (
	"hexagonal-go-grpc/internal/ports/secondary/db"
	"hexagonal-go-grpc/internal/ports/secondary/db/manager"
	"hexagonal-go-grpc/internal/ports/service"
)

// Service is contains functions for customer service.
type Service struct {
	trxMgr manager.ITransactionManager
	dbRepo db.ICustomerDbRepository
}

// New to create new customer service.
func New(trxMgr manager.ITransactionManager, dbRepo db.ICustomerDbRepository) service.ICustomerService {
	return &Service{
		trxMgr: trxMgr,
		dbRepo: dbRepo,
	}
}

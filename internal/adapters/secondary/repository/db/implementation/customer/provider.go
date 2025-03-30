package customer

import (
	"gorm.io/gorm"
	"hexagonal-go-grpc/internal/ports/secondary/db"
)

// DbRepository is contains functions for customer db.
type DbRepository struct {
	db *gorm.DB
}

// New to create new customer db.
func New(db *gorm.DB) db.ICustomerDbRepository {
	return &DbRepository{
		db: db,
	}
}

package customer

import (
	"hexagonal-go-grpc/internal/ports/primary/grpc"
	"hexagonal-go-grpc/internal/ports/service"
)

// Service is contains functions for customer service.
type GrpcHandler struct {
	service service.ICustomerService
}

// New to create new customer service.
func New(service service.ICustomerService) grpc.ICustomerGrpcHandler {
	return &GrpcHandler{
		service: service,
	}
}

package gin

import (
	"github.com/gin-gonic/gin"
	"hexagonal-go-grpc/internal/ports/primary/grpc"
	"hexagonal-go-grpc/proto/gen/api/v1/v1connect"
)

type Handler struct {
	customerHandler grpc.ICustomerGrpcHandler
}

type Parameters struct {
	CustomerHandler grpc.ICustomerGrpcHandler
}

func NewHandler(parameters Parameters) *Handler {
	return &Handler{
		customerHandler: parameters.CustomerHandler,
	}
}

func (handler *Handler) Register(r *gin.Engine) {
	// Create gRPC handler
	path, grpcHandler := v1connect.NewCustomerMethodHandler(handler.customerHandler)

	// Wrap gRPC handler
	r.Any(path+"/*any", gin.WrapH(grpcHandler))
}

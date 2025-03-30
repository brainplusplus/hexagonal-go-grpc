package echo

import (
	"github.com/labstack/echo/v4"
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

func (handler *Handler) Register(r *echo.Echo) {
	path, grpcHandler := v1connect.NewCustomerMethodHandler(handler.customerHandler)
	r.Any(path+"*", echo.WrapHandler(grpcHandler))
}

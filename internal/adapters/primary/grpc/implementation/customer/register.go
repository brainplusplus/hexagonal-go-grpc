package customer

import (
	"context"
	"github.com/bufbuild/connect-go"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	apiv1 "hexagonal-go-grpc/proto/gen/api/v1"
)

func (grpc GrpcHandler) Register(ctx context.Context, request *connect.Request[apiv1.CustomerRegisterRequest]) (*connect.Response[apiv1.CustomerRegisterResponse], error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "grpcHandler:customer:Register")
	defer span.End()

	params := entity.CustomerRegisterRequest{
		Name:   request.Msg.Name,
		Email:  request.Msg.Email,
		Phone:  request.Msg.Phone,
		Gender: request.Msg.Gender,
	}

	customer, statusCode, err := grpc.service.Register(ctx, params)

	if err != nil {
		res := &apiv1.CustomerRegisterResponse{
			StatusCode: int32(statusCode),
			Message:    err.Error(),
		}
		return connect.NewResponse(res), err
	}

	res := &apiv1.CustomerRegisterResponse{
		StatusCode: 200,
		Message:    "Successfully registered",
		Data: &apiv1.Customer{
			Id:        customer.ID,
			Name:      customer.Name,
			Email:     customer.Email,
			Phone:     customer.Phone,
			Gender:    customer.Gender,
			IsActive:  customer.IsActive,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
		},
	}

	return connect.NewResponse(res), err
}

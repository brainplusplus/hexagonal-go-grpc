package customer

import (
	"context"
	"github.com/bufbuild/connect-go"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	apiv1 "hexagonal-go-grpc/proto/gen/api/v1"
)

func (grpc GrpcHandler) GetById(ctx context.Context, request *connect.Request[apiv1.CustomerRequest]) (*connect.Response[apiv1.CustomerResponse], error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "grpcHandler:customer:GetById")
	defer span.End()

	params := entity.CustomerIdRequest{
		Id: request.Msg.Id,
	}

	customer, statusCode, err := grpc.service.GetById(ctx, params)

	if err != nil {
		res := &apiv1.CustomerResponse{
			StatusCode: int32(statusCode),
			Message:    err.Error(),
		}
		return connect.NewResponse(res), err
	}

	res := &apiv1.CustomerResponse{
		StatusCode: 200,
		Message:    "Successfully get customer",
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

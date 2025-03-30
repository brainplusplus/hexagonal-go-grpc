package customer

import (
	"context"
	"github.com/bufbuild/connect-go"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	apiv1 "hexagonal-go-grpc/proto/gen/api/v1"
)

func (grpc GrpcHandler) Update(ctx context.Context, request *connect.Request[apiv1.CustomerUpdateRequest]) (*connect.Response[apiv1.CustomerUpdateResponse], error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "grpcHandler:customer:Update")
	defer span.End()

	params := entity.CustomerUpdateRequest{
		Name:     request.Msg.Name,
		Email:    request.Msg.Email,
		Phone:    request.Msg.Phone,
		Gender:   request.Msg.Gender,
		IsActive: request.Msg.IsActive,
	}

	customer, statusCode, err := grpc.service.Update(ctx, params)

	if err != nil {
		res := &apiv1.CustomerUpdateResponse{
			StatusCode: int32(statusCode),
			Message:    err.Error(),
		}
		return connect.NewResponse(res), err
	}

	res := &apiv1.CustomerUpdateResponse{
		StatusCode: 200,
		Message:    "Successfully updated",
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

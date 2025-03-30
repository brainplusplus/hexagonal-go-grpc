package customer

import (
	"context"
	"github.com/bufbuild/connect-go"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	apiv1 "hexagonal-go-grpc/proto/gen/api/v1"
)

func (grpc GrpcHandler) GetList(ctx context.Context, request *connect.Request[apiv1.CustomerListRequest]) (*connect.Response[apiv1.CustomerListResponse], error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "grpcHandler:customer:GetList")
	defer span.End()

	params := entity.CustomerListRequest{
		Ids:    request.Msg.Ids,
		Search: request.Msg.Search,
	}

	customers, statusCode, err := grpc.service.GetList(ctx, params)

	if err != nil {
		res := &apiv1.CustomerListResponse{
			StatusCode: int32(statusCode),
			Message:    err.Error(),
			Data:       make([]*apiv1.Customer, 0),
		}
		return connect.NewResponse(res), err
	}

	data := make([]*apiv1.Customer, len(customers))
	for i, item := range customers {
		data[i] = &apiv1.Customer{
			Id:        item.ID,
			Name:      item.Name,
			Email:     item.Email,
			Phone:     item.Phone,
			Gender:    item.Gender,
			IsActive:  item.IsActive,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
	}

	res := &apiv1.CustomerListResponse{
		StatusCode: 200,
		Message:    "Successfully get list customer",
		Data:       data,
	}

	return connect.NewResponse(res), err
}

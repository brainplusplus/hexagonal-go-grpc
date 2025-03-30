package customer

import (
	"context"
	"github.com/bufbuild/connect-go"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	apiv1 "hexagonal-go-grpc/proto/gen/api/v1"
)

func (grpc GrpcHandler) ActivateById(ctx context.Context, request *connect.Request[apiv1.CustomerActivateRequest]) (*connect.Response[apiv1.CustomerActivateResponse], error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "grpcHandler:customer:ActivateById")
	defer span.End()

	params := entity.CustomerIdRequest{
		Id: request.Msg.Id,
	}
	_, statusCode, err := grpc.service.ActivateById(ctx, params)

	if err != nil {
		res := &apiv1.CustomerActivateResponse{
			StatusCode: int32(statusCode),
			Message:    err.Error(),
		}
		return connect.NewResponse(res), err
	}

	res := &apiv1.CustomerActivateResponse{
		StatusCode: 200,
		Message:    "Successfully activated",
	}

	return connect.NewResponse(res), err
}

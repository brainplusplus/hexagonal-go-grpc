package customer

import (
	"context"
	"github.com/bufbuild/connect-go"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	apiv1 "hexagonal-go-grpc/proto/gen/api/v1"
)

func (grpc GrpcHandler) DeactivateById(ctx context.Context, request *connect.Request[apiv1.CustomerDeactivateRequest]) (*connect.Response[apiv1.CustomerDeactivateResponse], error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "grpcHandler:customer:DeactivateById")
	defer span.End()

	params := entity.CustomerIdRequest{
		Id: request.Msg.Id,
	}

	_, statusCode, err := grpc.service.DeactivateById(ctx, params)

	if err != nil {
		res := &apiv1.CustomerDeactivateResponse{
			StatusCode: int32(statusCode),
			Message:    err.Error(),
		}
		return connect.NewResponse(res), err
	}

	res := &apiv1.CustomerDeactivateResponse{
		StatusCode: 200,
		Message:    "Successfully deactivated",
	}

	return connect.NewResponse(res), err
}

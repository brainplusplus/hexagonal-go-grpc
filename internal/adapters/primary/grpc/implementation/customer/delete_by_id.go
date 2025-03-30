package customer

import (
	"context"
	"github.com/bufbuild/connect-go"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	apiv1 "hexagonal-go-grpc/proto/gen/api/v1"
)

func (grpc GrpcHandler) DeleteById(ctx context.Context, request *connect.Request[apiv1.CustomerDeleteRequest]) (*connect.Response[apiv1.CustomerDeleteResponse], error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "grpcHandler:customer:DeleteById")
	defer span.End()

	params := entity.CustomerIdRequest{
		Id: request.Msg.Id,
	}

	statusCode, err := grpc.service.DeleteById(ctx, params)

	if err != nil {
		res := &apiv1.CustomerDeleteResponse{
			StatusCode: int32(statusCode),
			Message:    err.Error(),
		}
		return connect.NewResponse(res), err
	}

	res := &apiv1.CustomerDeleteResponse{
		StatusCode: 200,
		Message:    "Successfully deleted",
	}

	return connect.NewResponse(res), err
}

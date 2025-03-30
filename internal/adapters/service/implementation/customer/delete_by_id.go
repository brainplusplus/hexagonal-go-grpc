package customer

import (
	"context"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/service/entity"
)

func (service *Service) DeleteById(ctx context.Context, request entity.CustomerIdRequest) (statusCode int, err error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "service:customer:DeleteById")
	defer span.End()

	return service.dbRepo.DeleteById(ctx, request)
}

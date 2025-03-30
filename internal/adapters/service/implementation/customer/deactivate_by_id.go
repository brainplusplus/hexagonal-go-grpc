package customer

import (
	"context"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/service/entity"
)

func (service *Service) DeactivateById(ctx context.Context, request entity.CustomerIdRequest) (data entity.Customer, statusCode int, err error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "service:customer:DeactivateById")
	defer span.End()

	return service.dbRepo.DeactivateById(ctx, request)
}

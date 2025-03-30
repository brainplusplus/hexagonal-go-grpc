package customer

import (
	"context"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/service/entity"
)

func (service *Service) ActivateById(ctx context.Context, request entity.CustomerIdRequest) (data entity.Customer, statusCode int, err error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "service:customer:ActivateById")
	defer span.End()

	return service.dbRepo.ActivateById(ctx, request)
}

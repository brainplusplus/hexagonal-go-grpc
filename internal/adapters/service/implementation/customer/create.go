package customer

import (
	"context"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/service/entity"
)

func (service *Service) Create(ctx context.Context, request entity.CustomerCreateRequest) (data entity.Customer, statusCode int, err error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "service:customer:Create")
	defer span.End()

	return service.dbRepo.Create(ctx, request)
}

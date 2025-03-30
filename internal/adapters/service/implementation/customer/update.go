package customer

import (
	"context"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/service/entity"
)

func (service *Service) Update(ctx context.Context, request entity.CustomerUpdateRequest) (data entity.Customer, statusCode int, err error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "service:customer:Update")
	defer span.End()

	return service.dbRepo.Update(ctx, request)
}

package customer

import (
	"context"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/service/entity"
)

func (service *Service) GetById(ctx context.Context, request entity.CustomerIdRequest) (data entity.Customer, statusCode int, err error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "service:customer:GetById")
	defer span.End()

	return service.dbRepo.GetById(ctx, request)
}

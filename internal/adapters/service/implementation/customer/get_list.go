package customer

import (
	"context"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/service/entity"
)

func (service *Service) GetList(ctx context.Context, request entity.CustomerListRequest) ([]entity.Customer, int, error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "service:customer:GetList")
	defer span.End()

	return service.dbRepo.GetList(ctx, request)
}

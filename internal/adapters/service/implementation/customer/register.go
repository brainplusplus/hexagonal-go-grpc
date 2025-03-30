package customer

import (
	"context"
	"hexagonal-go-grpc/internal/adapters/errors"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	"net/http"
)

func (service *Service) Register(ctx context.Context, request entity.CustomerRegisterRequest) (entity.Customer, int, error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "service:customer:Register")
	defer span.End()

	results, statusCode, err := service.trxMgr.WithTransaction(func(ctx context.Context, args ...any) ([]any, int, error) {
		request := args[0].(entity.CustomerRegisterRequest)
		var customer entity.Customer
		_, statusCode, err := service.dbRepo.GetByEmailOrPhone(ctx, entity.CustomerEmailOrPhoneRequest{
			Email: request.Email,
			Phone: request.Phone,
		})

		if err != nil {
			if statusCode == http.StatusNotFound {
				customer, statusCode, err = service.dbRepo.Create(ctx, entity.CustomerCreateRequest{
					Name:     request.Name,
					Email:    request.Email,
					Phone:    request.Phone,
					Gender:   request.Gender,
					IsActive: false,
				})
				return []any{customer}, statusCode, err
			} else {
				return nil, http.StatusBadRequest, errors.ErrFailedCreateCustomer
			}
		}
		return nil, http.StatusBadRequest, errors.ErrFailedCreateCustomer
	},
		ctx,
		request)

	if results != nil {
		return results[0].(entity.Customer), statusCode, err
	} else {
		return entity.Customer{}, statusCode, err
	}
}

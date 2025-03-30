package customer

import (
	"context"
	"hexagonal-go-grpc/internal/adapters/errors"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/secondary/repository/db/model"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	"hexagonal-go-grpc/internal/adapters/util"
	"log/slog"
	"net/http"
)

func (dbRepository *DbRepository) Create(ctx context.Context, customer entity.CustomerCreateRequest) (data entity.Customer, statusCode int, err error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "secondary:db:customer:Create")
	defer span.End()

	customerDb := model.CustomerCreateEntityToModel(customer)
	db := util.GetDBFromContext(ctx, dbRepository.db)

	err = db.Create(&customerDb).Error
	if err != nil {
		slog.ErrorContext(ctx, "failed to create customer", slog.Any("err ", err))
		return entity.Customer{}, http.StatusInternalServerError, errors.Wrap(ctx, errors.ErrFailedCreateCustomer, err)
	}
	data = customerDb.ToEntity()
	return data, http.StatusOK, nil
}

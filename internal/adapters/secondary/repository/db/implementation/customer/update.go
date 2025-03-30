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

func (dbRepository *DbRepository) Update(ctx context.Context, request entity.CustomerUpdateRequest) (data entity.Customer, statusCode int, err error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "secondary:db:customer:Update")
	defer span.End()

	customerDb := model.CustomerUpdateEntityToModel(request)
	customerDb.ID = request.ID

	db := util.GetDBFromContext(ctx, dbRepository.db)
	err = db.Model(&entity.Customer{}).Where("id = ?", customerDb.ID).Updates(&customerDb).Error
	if err != nil {
		slog.ErrorContext(ctx, "failed to update customer", slog.Any("err ", err))
		return entity.Customer{}, http.StatusInternalServerError, errors.Wrap(ctx, errors.ErrFailedUpdateCustomer, err)
	}
	data = customerDb.ToEntity()
	return data, http.StatusOK, nil
}

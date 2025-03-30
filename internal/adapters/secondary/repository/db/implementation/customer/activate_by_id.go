package customer

import (
	"context"
	"gorm.io/gorm"
	"hexagonal-go-grpc/internal/adapters/errors"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/secondary/repository/db/model"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	"hexagonal-go-grpc/internal/adapters/util"
	"log/slog"
	"net/http"
)

func (dbRepository *DbRepository) ActivateById(ctx context.Context, request entity.CustomerIdRequest) (data entity.Customer, statusCode int, err error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "secondary:db:customer:ActivateById")
	defer span.End()

	var customerDb model.Customer

	db := util.GetDBFromContext(ctx, dbRepository.db)
	err = db.Where("id = ?", request.Id).First(&customerDb).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		slog.ErrorContext(ctx, errors.ErrFailedGetCustomer.Error(), slog.Any("err ", err))
		return entity.Customer{}, http.StatusNotFound, errors.ErrFailedGetCustomer
	} else if err != nil {
		slog.ErrorContext(ctx, errors.ErrInternalDB.Error(), slog.Any("err ", err))
		return entity.Customer{}, http.StatusInternalServerError, errors.ErrInternalDB
	}

	customerDb.IsActive = true
	err = db.Save(&customerDb).Error
	if err != nil {
		slog.ErrorContext(ctx, errors.ErrFailedActivateCustomer.Error(), slog.Any("err ", err))
		return entity.Customer{}, http.StatusInternalServerError, errors.ErrFailedActivateCustomer
	}
	data = customerDb.ToEntity()
	return data, http.StatusOK, nil
}

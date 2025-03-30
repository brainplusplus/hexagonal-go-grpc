package customer

import (
	"context"
	"gorm.io/gorm"
	"hexagonal-go-grpc/internal/adapters/errors"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/secondary/repository/db/model"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	"hexagonal-go-grpc/internal/adapters/util"
	"net/http"
)

func (dbRepository *DbRepository) DeleteById(ctx context.Context, request entity.CustomerIdRequest) (statusCode int, err error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "secondary:db:customer:DeleteById")
	defer span.End()

	var customerDb model.Customer

	db := util.GetDBFromContext(ctx, dbRepository.db)

	err = db.Where("id = ?", request.Id).First(&customerDb).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound, err
	} else if err != nil {
		return http.StatusInternalServerError, err
	}

	err = db.Delete(&customerDb).Error
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

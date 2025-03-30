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

func (dbRepository *DbRepository) GetByEmailOrPhone(ctx context.Context, request entity.CustomerEmailOrPhoneRequest) (data entity.Customer, statusCode int, err error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "secondary:rdbms:customer:GetByEmailOrPhone")
	defer span.End()

	var customerDb model.Customer

	db := util.GetDBFromContext(ctx, dbRepository.db)

	err = db.Where("email = ? OR phone = ?", request.Email, request.Phone).First(&customerDb).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return data, http.StatusNotFound, err
	} else if err != nil {
		return data, http.StatusInternalServerError, err
	}
	data = customerDb.ToEntity()

	return data, http.StatusOK, nil
}

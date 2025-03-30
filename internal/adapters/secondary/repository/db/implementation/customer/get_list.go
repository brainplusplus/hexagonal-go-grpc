package customer

import (
	"context"
	"fmt"
	"hexagonal-go-grpc/internal/adapters/errors"
	infraObsrv "hexagonal-go-grpc/internal/adapters/infrastructure/observability"
	"hexagonal-go-grpc/internal/adapters/secondary/repository/db/model"
	"hexagonal-go-grpc/internal/adapters/service/entity"
	"hexagonal-go-grpc/internal/adapters/util"
	"log/slog"
	"net/http"
)

func (dbRepository *DbRepository) GetList(ctx context.Context, request entity.CustomerListRequest) ([]entity.Customer, int, error) {
	ctx, span := infraObsrv.Tracer().Start(ctx, "secondary:db:customer:GetList")
	defer span.End()

	var customersDb []model.Customer
	db := util.GetDBFromContext(ctx, dbRepository.db)

	query := db.Model(&entity.Customer{})
	if len(request.Ids) > 0 {
		query = query.Where("id IN (?)", request.Ids)
	}
	if len(request.Search) > 0 {
		search := fmt.Sprintf("%%%s%%", request.Search)
		query = query.Where("(name LIKE ? OR email LIKE ? OR phone LIKE ?)", search, search, search)
	}

	err := query.Find(&customersDb).Error
	if err != nil {
		slog.ErrorContext(ctx, "failed to get list customer", slog.Any("err ", err))
		return nil, http.StatusInternalServerError, errors.ErrFailedGetListCustomer
	}
	customers := model.CustomerModelListToEntityList(customersDb)
	return customers, http.StatusOK, nil
}

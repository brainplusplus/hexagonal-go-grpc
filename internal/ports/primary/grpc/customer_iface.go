package grpc

import (
	"context"
	"github.com/bufbuild/connect-go"
	apiv1 "hexagonal-go-grpc/proto/gen/api/v1"
)

type ICustomerGrpcHandler interface {
	GetList(ctx context.Context, request *connect.Request[apiv1.CustomerListRequest]) (*connect.Response[apiv1.CustomerListResponse], error)
	GetById(ctx context.Context, request *connect.Request[apiv1.CustomerRequest]) (*connect.Response[apiv1.CustomerResponse], error)
	Create(ctx context.Context, request *connect.Request[apiv1.CustomerCreateRequest]) (*connect.Response[apiv1.CustomerCreateResponse], error)
	Update(ctx context.Context, request *connect.Request[apiv1.CustomerUpdateRequest]) (*connect.Response[apiv1.CustomerUpdateResponse], error)
	DeleteById(ctx context.Context, request *connect.Request[apiv1.CustomerDeleteRequest]) (*connect.Response[apiv1.CustomerDeleteResponse], error)
	ActivateById(ctx context.Context, request *connect.Request[apiv1.CustomerActivateRequest]) (*connect.Response[apiv1.CustomerActivateResponse], error)
	DeactivateById(ctx context.Context, request *connect.Request[apiv1.CustomerDeactivateRequest]) (*connect.Response[apiv1.CustomerDeactivateResponse], error)
	Register(ctx context.Context, request *connect.Request[apiv1.CustomerRegisterRequest]) (*connect.Response[apiv1.CustomerRegisterResponse], error)
}

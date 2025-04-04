// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/api/v1/customer.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "hexagonal-go-grpc/proto/gen/api/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// CustomerMethodName is the fully-qualified name of the CustomerMethod service.
	CustomerMethodName = "api.v1.CustomerMethod"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CustomerMethodGetListProcedure is the fully-qualified name of the CustomerMethod's GetList RPC.
	CustomerMethodGetListProcedure = "/api.v1.CustomerMethod/GetList"
	// CustomerMethodGetByIdProcedure is the fully-qualified name of the CustomerMethod's GetById RPC.
	CustomerMethodGetByIdProcedure = "/api.v1.CustomerMethod/GetById"
	// CustomerMethodCreateProcedure is the fully-qualified name of the CustomerMethod's Create RPC.
	CustomerMethodCreateProcedure = "/api.v1.CustomerMethod/Create"
	// CustomerMethodUpdateProcedure is the fully-qualified name of the CustomerMethod's Update RPC.
	CustomerMethodUpdateProcedure = "/api.v1.CustomerMethod/Update"
	// CustomerMethodDeleteByIdProcedure is the fully-qualified name of the CustomerMethod's DeleteById
	// RPC.
	CustomerMethodDeleteByIdProcedure = "/api.v1.CustomerMethod/DeleteById"
	// CustomerMethodActivateByIdProcedure is the fully-qualified name of the CustomerMethod's
	// ActivateById RPC.
	CustomerMethodActivateByIdProcedure = "/api.v1.CustomerMethod/ActivateById"
	// CustomerMethodDeactivateByIdProcedure is the fully-qualified name of the CustomerMethod's
	// DeactivateById RPC.
	CustomerMethodDeactivateByIdProcedure = "/api.v1.CustomerMethod/DeactivateById"
	// CustomerMethodRegisterProcedure is the fully-qualified name of the CustomerMethod's Register RPC.
	CustomerMethodRegisterProcedure = "/api.v1.CustomerMethod/Register"
)

// CustomerMethodClient is a client for the api.v1.CustomerMethod service.
type CustomerMethodClient interface {
	GetList(context.Context, *connect_go.Request[v1.CustomerListRequest]) (*connect_go.Response[v1.CustomerListResponse], error)
	GetById(context.Context, *connect_go.Request[v1.CustomerRequest]) (*connect_go.Response[v1.CustomerResponse], error)
	Create(context.Context, *connect_go.Request[v1.CustomerCreateRequest]) (*connect_go.Response[v1.CustomerCreateResponse], error)
	Update(context.Context, *connect_go.Request[v1.CustomerUpdateRequest]) (*connect_go.Response[v1.CustomerUpdateResponse], error)
	DeleteById(context.Context, *connect_go.Request[v1.CustomerDeleteRequest]) (*connect_go.Response[v1.CustomerDeleteResponse], error)
	ActivateById(context.Context, *connect_go.Request[v1.CustomerActivateRequest]) (*connect_go.Response[v1.CustomerActivateResponse], error)
	DeactivateById(context.Context, *connect_go.Request[v1.CustomerDeactivateRequest]) (*connect_go.Response[v1.CustomerDeactivateResponse], error)
	Register(context.Context, *connect_go.Request[v1.CustomerRegisterRequest]) (*connect_go.Response[v1.CustomerRegisterResponse], error)
}

// NewCustomerMethodClient constructs a client for the api.v1.CustomerMethod service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCustomerMethodClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CustomerMethodClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &customerMethodClient{
		getList: connect_go.NewClient[v1.CustomerListRequest, v1.CustomerListResponse](
			httpClient,
			baseURL+CustomerMethodGetListProcedure,
			opts...,
		),
		getById: connect_go.NewClient[v1.CustomerRequest, v1.CustomerResponse](
			httpClient,
			baseURL+CustomerMethodGetByIdProcedure,
			opts...,
		),
		create: connect_go.NewClient[v1.CustomerCreateRequest, v1.CustomerCreateResponse](
			httpClient,
			baseURL+CustomerMethodCreateProcedure,
			opts...,
		),
		update: connect_go.NewClient[v1.CustomerUpdateRequest, v1.CustomerUpdateResponse](
			httpClient,
			baseURL+CustomerMethodUpdateProcedure,
			opts...,
		),
		deleteById: connect_go.NewClient[v1.CustomerDeleteRequest, v1.CustomerDeleteResponse](
			httpClient,
			baseURL+CustomerMethodDeleteByIdProcedure,
			opts...,
		),
		activateById: connect_go.NewClient[v1.CustomerActivateRequest, v1.CustomerActivateResponse](
			httpClient,
			baseURL+CustomerMethodActivateByIdProcedure,
			opts...,
		),
		deactivateById: connect_go.NewClient[v1.CustomerDeactivateRequest, v1.CustomerDeactivateResponse](
			httpClient,
			baseURL+CustomerMethodDeactivateByIdProcedure,
			opts...,
		),
		register: connect_go.NewClient[v1.CustomerRegisterRequest, v1.CustomerRegisterResponse](
			httpClient,
			baseURL+CustomerMethodRegisterProcedure,
			opts...,
		),
	}
}

// customerMethodClient implements CustomerMethodClient.
type customerMethodClient struct {
	getList        *connect_go.Client[v1.CustomerListRequest, v1.CustomerListResponse]
	getById        *connect_go.Client[v1.CustomerRequest, v1.CustomerResponse]
	create         *connect_go.Client[v1.CustomerCreateRequest, v1.CustomerCreateResponse]
	update         *connect_go.Client[v1.CustomerUpdateRequest, v1.CustomerUpdateResponse]
	deleteById     *connect_go.Client[v1.CustomerDeleteRequest, v1.CustomerDeleteResponse]
	activateById   *connect_go.Client[v1.CustomerActivateRequest, v1.CustomerActivateResponse]
	deactivateById *connect_go.Client[v1.CustomerDeactivateRequest, v1.CustomerDeactivateResponse]
	register       *connect_go.Client[v1.CustomerRegisterRequest, v1.CustomerRegisterResponse]
}

// GetList calls api.v1.CustomerMethod.GetList.
func (c *customerMethodClient) GetList(ctx context.Context, req *connect_go.Request[v1.CustomerListRequest]) (*connect_go.Response[v1.CustomerListResponse], error) {
	return c.getList.CallUnary(ctx, req)
}

// GetById calls api.v1.CustomerMethod.GetById.
func (c *customerMethodClient) GetById(ctx context.Context, req *connect_go.Request[v1.CustomerRequest]) (*connect_go.Response[v1.CustomerResponse], error) {
	return c.getById.CallUnary(ctx, req)
}

// Create calls api.v1.CustomerMethod.Create.
func (c *customerMethodClient) Create(ctx context.Context, req *connect_go.Request[v1.CustomerCreateRequest]) (*connect_go.Response[v1.CustomerCreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Update calls api.v1.CustomerMethod.Update.
func (c *customerMethodClient) Update(ctx context.Context, req *connect_go.Request[v1.CustomerUpdateRequest]) (*connect_go.Response[v1.CustomerUpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// DeleteById calls api.v1.CustomerMethod.DeleteById.
func (c *customerMethodClient) DeleteById(ctx context.Context, req *connect_go.Request[v1.CustomerDeleteRequest]) (*connect_go.Response[v1.CustomerDeleteResponse], error) {
	return c.deleteById.CallUnary(ctx, req)
}

// ActivateById calls api.v1.CustomerMethod.ActivateById.
func (c *customerMethodClient) ActivateById(ctx context.Context, req *connect_go.Request[v1.CustomerActivateRequest]) (*connect_go.Response[v1.CustomerActivateResponse], error) {
	return c.activateById.CallUnary(ctx, req)
}

// DeactivateById calls api.v1.CustomerMethod.DeactivateById.
func (c *customerMethodClient) DeactivateById(ctx context.Context, req *connect_go.Request[v1.CustomerDeactivateRequest]) (*connect_go.Response[v1.CustomerDeactivateResponse], error) {
	return c.deactivateById.CallUnary(ctx, req)
}

// Register calls api.v1.CustomerMethod.Register.
func (c *customerMethodClient) Register(ctx context.Context, req *connect_go.Request[v1.CustomerRegisterRequest]) (*connect_go.Response[v1.CustomerRegisterResponse], error) {
	return c.register.CallUnary(ctx, req)
}

// CustomerMethodHandler is an implementation of the api.v1.CustomerMethod service.
type CustomerMethodHandler interface {
	GetList(context.Context, *connect_go.Request[v1.CustomerListRequest]) (*connect_go.Response[v1.CustomerListResponse], error)
	GetById(context.Context, *connect_go.Request[v1.CustomerRequest]) (*connect_go.Response[v1.CustomerResponse], error)
	Create(context.Context, *connect_go.Request[v1.CustomerCreateRequest]) (*connect_go.Response[v1.CustomerCreateResponse], error)
	Update(context.Context, *connect_go.Request[v1.CustomerUpdateRequest]) (*connect_go.Response[v1.CustomerUpdateResponse], error)
	DeleteById(context.Context, *connect_go.Request[v1.CustomerDeleteRequest]) (*connect_go.Response[v1.CustomerDeleteResponse], error)
	ActivateById(context.Context, *connect_go.Request[v1.CustomerActivateRequest]) (*connect_go.Response[v1.CustomerActivateResponse], error)
	DeactivateById(context.Context, *connect_go.Request[v1.CustomerDeactivateRequest]) (*connect_go.Response[v1.CustomerDeactivateResponse], error)
	Register(context.Context, *connect_go.Request[v1.CustomerRegisterRequest]) (*connect_go.Response[v1.CustomerRegisterResponse], error)
}

// NewCustomerMethodHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCustomerMethodHandler(svc CustomerMethodHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	customerMethodGetListHandler := connect_go.NewUnaryHandler(
		CustomerMethodGetListProcedure,
		svc.GetList,
		opts...,
	)
	customerMethodGetByIdHandler := connect_go.NewUnaryHandler(
		CustomerMethodGetByIdProcedure,
		svc.GetById,
		opts...,
	)
	customerMethodCreateHandler := connect_go.NewUnaryHandler(
		CustomerMethodCreateProcedure,
		svc.Create,
		opts...,
	)
	customerMethodUpdateHandler := connect_go.NewUnaryHandler(
		CustomerMethodUpdateProcedure,
		svc.Update,
		opts...,
	)
	customerMethodDeleteByIdHandler := connect_go.NewUnaryHandler(
		CustomerMethodDeleteByIdProcedure,
		svc.DeleteById,
		opts...,
	)
	customerMethodActivateByIdHandler := connect_go.NewUnaryHandler(
		CustomerMethodActivateByIdProcedure,
		svc.ActivateById,
		opts...,
	)
	customerMethodDeactivateByIdHandler := connect_go.NewUnaryHandler(
		CustomerMethodDeactivateByIdProcedure,
		svc.DeactivateById,
		opts...,
	)
	customerMethodRegisterHandler := connect_go.NewUnaryHandler(
		CustomerMethodRegisterProcedure,
		svc.Register,
		opts...,
	)
	return "/api.v1.CustomerMethod/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CustomerMethodGetListProcedure:
			customerMethodGetListHandler.ServeHTTP(w, r)
		case CustomerMethodGetByIdProcedure:
			customerMethodGetByIdHandler.ServeHTTP(w, r)
		case CustomerMethodCreateProcedure:
			customerMethodCreateHandler.ServeHTTP(w, r)
		case CustomerMethodUpdateProcedure:
			customerMethodUpdateHandler.ServeHTTP(w, r)
		case CustomerMethodDeleteByIdProcedure:
			customerMethodDeleteByIdHandler.ServeHTTP(w, r)
		case CustomerMethodActivateByIdProcedure:
			customerMethodActivateByIdHandler.ServeHTTP(w, r)
		case CustomerMethodDeactivateByIdProcedure:
			customerMethodDeactivateByIdHandler.ServeHTTP(w, r)
		case CustomerMethodRegisterProcedure:
			customerMethodRegisterHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCustomerMethodHandler returns CodeUnimplemented from all methods.
type UnimplementedCustomerMethodHandler struct{}

func (UnimplementedCustomerMethodHandler) GetList(context.Context, *connect_go.Request[v1.CustomerListRequest]) (*connect_go.Response[v1.CustomerListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.CustomerMethod.GetList is not implemented"))
}

func (UnimplementedCustomerMethodHandler) GetById(context.Context, *connect_go.Request[v1.CustomerRequest]) (*connect_go.Response[v1.CustomerResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.CustomerMethod.GetById is not implemented"))
}

func (UnimplementedCustomerMethodHandler) Create(context.Context, *connect_go.Request[v1.CustomerCreateRequest]) (*connect_go.Response[v1.CustomerCreateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.CustomerMethod.Create is not implemented"))
}

func (UnimplementedCustomerMethodHandler) Update(context.Context, *connect_go.Request[v1.CustomerUpdateRequest]) (*connect_go.Response[v1.CustomerUpdateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.CustomerMethod.Update is not implemented"))
}

func (UnimplementedCustomerMethodHandler) DeleteById(context.Context, *connect_go.Request[v1.CustomerDeleteRequest]) (*connect_go.Response[v1.CustomerDeleteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.CustomerMethod.DeleteById is not implemented"))
}

func (UnimplementedCustomerMethodHandler) ActivateById(context.Context, *connect_go.Request[v1.CustomerActivateRequest]) (*connect_go.Response[v1.CustomerActivateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.CustomerMethod.ActivateById is not implemented"))
}

func (UnimplementedCustomerMethodHandler) DeactivateById(context.Context, *connect_go.Request[v1.CustomerDeactivateRequest]) (*connect_go.Response[v1.CustomerDeactivateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.CustomerMethod.DeactivateById is not implemented"))
}

func (UnimplementedCustomerMethodHandler) Register(context.Context, *connect_go.Request[v1.CustomerRegisterRequest]) (*connect_go.Response[v1.CustomerRegisterResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.CustomerMethod.Register is not implemented"))
}

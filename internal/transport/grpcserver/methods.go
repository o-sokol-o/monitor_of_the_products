package grpcserver

import (
	"context"
	"fmt"

	"github.com/o-sokol-o/as/gen/grpc_product"
	"github.com/o-sokol-o/as/pkg/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IAppServices interface {
	FetchExternProducts(ctx context.Context, url string) (int, int, error)
	ListExternProducts(ctx context.Context, paging domain.Paging, sorting domain.Sorting) (domain.Products, error)

	InsertLog(ctx context.Context, log domain.LogItem) error
	InsertProducts(ctx context.Context, products domain.Products) error
	UpdateProducts(ctx context.Context, products domain.Products) error
	GetProducts(ctx context.Context) (domain.Products, error)
}

// Реализация методов gRPC сервера: Интерфейс GrpcProductsServiceServer ищем в сгенерированом файле.
// Имя интерфейса "GrpcProductsServiceServer" образуется из вбитого имени сервиса + Server
type ServerMethods struct {
	services IAppServices
	grpc_product.UnimplementedGrpcProductsServiceServer
}

func NewServerMethods(services IAppServices) *ServerMethods {
	return &ServerMethods{
		services: services,
	}
}

func (h *ServerMethods) Fetch(ctx context.Context, req *grpc_product.RequestFetch) (*grpc_product.ResponseFetch, error) {
	err := h.services.InsertLog(ctx, domain.LogItem{RequestType: "FETCH", Input: req.GetUrl()})
	if err != nil {
		return &grpc_product.ResponseFetch{NewProductCount: 0, UpdateProductCount: 0}, err
	}

	// Ping ?
	if req.Url == "" {
		fmt.Printf("msg receive: Ping\n")
		return &grpc_product.ResponseFetch{NewProductCount: 0, UpdateProductCount: 0}, nil
	}

	fmt.Printf("msg receive: Fetch: %v\n", req)
	newProd, updateProd, err := h.services.FetchExternProducts(ctx, req.Url)
	if err != nil {
		fmt.Printf("error: %+v\n", err)
	}

	return &grpc_product.ResponseFetch{NewProductCount: int32(newProd), UpdateProductCount: int32(updateProd)}, err
}

func (h *ServerMethods) List(ctx context.Context, req *grpc_product.RequestList) (*grpc_product.ResponseList, error) {
	fmt.Printf("msg receive: List: %v\n", req)
	err := h.services.InsertLog(ctx, domain.LogItem{RequestType: "LIST", Input: fmt.Sprintf("%v", req)})
	if err != nil {
		return nil, err
	}

	products, err := h.services.ListExternProducts(ctx,
		domain.Paging{NumPage: req.Paging.NumPage, ProductPerPage: req.Paging.ProductPerPage},
		domain.Sorting{SortField: req.Sorting.SortField, Direction: req.Sorting.Direction})
	if err != nil {
		fmt.Printf("error: %+v\n", err)
		return nil, err
	}

	var gp []*grpc_product.Product
	for _, pr := range products {
		gp = append(gp, &grpc_product.Product{
			Name:        pr.Name,
			Cost:        pr.Cost,
			ChangeCount: pr.ChangeCount,
			UpdateAt:    &timestamppb.Timestamp{Seconds: int64(pr.UpdateAt.Second()), Nanos: int32(pr.UpdateAt.Nanosecond())}})
	}

	return &grpc_product.ResponseList{Products: gp}, err
}

package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/o-sokol-o/as/gen/grpc_product"
	"github.com/o-sokol-o/as/pkg/domain"
)

type GrpcClient struct {
	client grpc_product.GrpcProductsServiceClient
	conn   *grpc.ClientConn
	msg_id int64
}

func New(target string) (*GrpcClient, error) {
	// Set up a connection to the server.
	// Target as "localhost:9000"
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	gc := GrpcClient{
		client: grpc_product.NewGrpcProductsServiceClient(conn),
		conn:   conn,
	}

	// Send Ping
	err = gc.Fetch("")
	if err != nil {
		return nil, errors.New("grpc ping error: " + err.Error())
	}

	return &gc, nil
}

func (c *GrpcClient) Fetch(url string) error {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	msg := grpc_product.RequestFetch{Url: url}
	r, err := c.client.Fetch(ctx, &msg)
	if err != nil {
		return err
	}
	fmt.Printf("Fetch %d sended. Response: New product = %d, Update product = %d\n", c.msg_id, r.NewProductCount, r.UpdateProductCount)
	c.msg_id++
	return nil
}

func (c *GrpcClient) List(paging domain.Paging, sorting domain.Sorting) (domain.Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pg := grpc_product.Paging{
		NumPage:        paging.NumPage,
		ProductPerPage: paging.ProductPerPage,
	}

	srt := grpc_product.Sorting{
		SortField: sorting.SortField,
		Direction: sorting.Direction,
	}

	msg := grpc_product.RequestList{
		Paging:  &pg,
		Sorting: &srt,
	}

	r, err := c.client.List(ctx, &msg)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Request list %d sended. Response:\n", c.msg_id)
	c.msg_id++

	var products domain.Products
	for _, pr := range r.Products {
		products = append(products, domain.Product{
			Name:        pr.Name,
			Cost:        pr.Cost,
			ChangeCount: pr.ChangeCount,
			UpdateAt:    pr.UpdateAt.AsTime(),
		})
	}

	return products, nil
}

func (c *GrpcClient) Close() {
	c.conn.Close()
}

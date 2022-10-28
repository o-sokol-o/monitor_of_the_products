package grpcserver

import (
	"fmt"
	"net"

	"github.com/o-sokol-o/as/gen/grpc_product"
	"google.golang.org/grpc"
)

type Server struct {
	grpcSrv        *grpc.Server
	grpcLogsServer grpc_product.GrpcProductsServiceServer
}

func NewGrpcServer(service IAppServices) *Server {
	return &Server{
		grpcSrv:        grpc.NewServer(),
		grpcLogsServer: NewServerMethods(service),
	}
}

func (s *Server) Run(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	grpc_product.RegisterGrpcProductsServiceServer(s.grpcSrv, s.grpcLogsServer)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}

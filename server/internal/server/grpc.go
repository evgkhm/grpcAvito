package server

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"grpcAvito/proto"
	"grpcAvito/server/internal/service"
	"net"
)

/*var (
	port = flag.Int("port", 50051, "The server port")
)*/

func NewGRPCServer(service *service.Service, logger *logrus.Logger) (*grpc.Server, net.Listener) {
	//flag.Parse()
	/*lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GRPC.Port))
	if err != nil {
		logger.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterServerServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}

	return s, lis*/
	logger.Info("new grpc server")
	log := logrus.NewEntry(logger)
	var opts = []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_logrus.StreamServerInterceptor(log),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(log),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	}

	srv := grpc.NewServer(opts...)
	proto.RegisterServerServer(srv, service)

	return srv
}

package grpc

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"grpcAvito/internal/service/wallet"
	"grpcAvito/internal/service/wallet/spec"
)

type Conf struct {
	Port     string
	HostPort string
}

var Config *Conf

func (h Conf) Init() {
	Config = &Conf{
		Port:     viper.GetString("grpc.port"),
		HostPort: viper.GetString("grpc.host_port"),
	}
}

func New(services *wallet.Service, logger *logrus.Logger) *grpc.Server {
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
	spec.RegisterWalletAppServiceServer(srv, services)

	return srv
}

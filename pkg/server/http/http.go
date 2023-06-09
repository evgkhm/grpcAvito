package http

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpcAvito/internal/service/spec"
)

func New(endpoint string, log *logrus.Logger) *runtime.ServeMux {
	log.Info("new http server")
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := spec.RegisterWalletAppServiceHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	if err != nil {
		log.Panic(err)
	}

	return mux
}

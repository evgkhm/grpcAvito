package http

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpcAvito/internal/service/wallet/spec"
	"time"
)

var Config *Conf

type Conf struct {
	Port     string
	HostPort string
	LogPath  string
	Duration time.Duration
}

func (h Conf) Init() {
	Config = &Conf{
		Port:     viper.GetString("server.port"),
		HostPort: viper.GetString("server.host_port"),
		LogPath:  viper.GetString("server.log_path"),
		Duration: viper.GetDuration("server.duration"),
	}
}

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

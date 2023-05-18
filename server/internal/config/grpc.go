package config

import "github.com/spf13/viper"

type GRPCConfig struct {
	Port     string
	HostPort string
}

var GRPC *GRPCConfig

func (h GRPCConfig) Init() {
	GRPC = &GRPCConfig{
		Port:     viper.GetString("grpc.port"),
		HostPort: viper.GetString("grpc.host_port"),
	}
}

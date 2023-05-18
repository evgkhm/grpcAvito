package config

import "github.com/spf13/viper"

type GRPCConfig struct {
	Port string
}

var GRPC *GRPCConfig

func (h GRPCConfig) Init() {
	GRPC = &GRPCConfig{
		Port: viper.GetString("grpc.port"),
	}
}

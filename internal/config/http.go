package config

import (
	"github.com/spf13/viper"
)

type HttpConfig struct {
	Port     string
	HostPort string
	LogPath  string
}

var HTTP *HttpConfig

func (h HttpConfig) Init() {
	HTTP = &HttpConfig{
		Port:     viper.GetString("server.port"),
		HostPort: viper.GetString("server.host_port"),
		LogPath:  viper.GetString("server.log_path"),
	}
}

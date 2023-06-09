package config

import (
	"github.com/spf13/viper"
	"time"
)

type HTTPConfig struct {
	Port     string
	HostPort string
	LogPath  string
	Duration time.Duration
}

var HTTP *HTTPConfig

func (h HTTPConfig) Init() {
	HTTP = &HTTPConfig{
		Port:     viper.GetString("server.port"),
		HostPort: viper.GetString("server.host_port"),
		LogPath:  viper.GetString("server.log_path"),
		Duration: viper.GetDuration("server.duration"),
	}
}

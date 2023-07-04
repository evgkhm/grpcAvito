package main

import (
	"grpcAvito/internal/app"
	config "grpcAvito/internal/config"
)

func init() {
	config.InitAll([]config.Config{
		config.PostgresConfig{},
		config.HTTPConfig{},
		config.GRPCConfig{},
	})
}

func main() {
	app.Run()
}

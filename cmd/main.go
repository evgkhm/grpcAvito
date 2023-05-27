package main

import (
	"context"

	"golang.org/x/sync/errgroup"

	config "grpcAvito/internal/config"
	"grpcAvito/internal/repository/postgres"
	"grpcAvito/internal/server"
	"grpcAvito/internal/service"
	"grpcAvito/internal/usecase"
	"grpcAvito/pkg/logging"
	"net"
	"net/http"
)

func init() {
	config.InitAll([]config.Config{
		config.PostgresConfig{},
		config.HttpConfig{},
		config.GRPCConfig{},
	})
}

func main() {
	log := logging.GetLogger()
	postgresDB, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to initialize database connection: %s", err.Error())
	}

	postgresRepository := postgres.New(postgresDB, log)

	useCases := usecase.New(postgresRepository, log, postgresDB)

	service := service.New(useCases, log)

	grpcServer := server.NewGRPCServer(service, log)

	listen, err := net.Listen("tcp", config.GRPC.HostPort)
	if err != nil {
		log.Panic(err)
	}

	mux := server.NewHTTPServer(config.GRPC.Port, log)

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() (err error) {
		return grpcServer.Serve(listen)
	})
	g.Go(func() (err error) {
		return http.ListenAndServe(config.HTTP.HostPort, mux)
	})

	err = g.Wait()
	if err != nil {
		log.Panic(err)
	}
}

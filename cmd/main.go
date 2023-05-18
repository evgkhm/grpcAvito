package main

import (
	"context"

	"golang.org/x/sync/errgroup"

	config2 "grpcAvito/internal/config"
	"grpcAvito/internal/repository/postgres"
	"grpcAvito/internal/server"
	"grpcAvito/internal/service"
	"grpcAvito/internal/usecase"
	"grpcAvito/pkg/logging"
	"net"
	"net/http"
)

func init() {
	config2.InitAll([]config2.Config{
		config2.PostgresConfig{},
		config2.HttpConfig{},
		config2.GRPCConfig{},
	})
}

func main() {
	log := logging.GetLogger()
	postgresDB, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to initialize database connection: %s", err.Error())
	}

	postgresRepository := postgres.NewRepositoryPostgres(postgresDB, log)

	useCases := usecase.NewUseCase(postgresRepository, log)

	service := service.NewService(useCases, log)

	grpcServer := server.NewGRPCServer(service, log)

	listen, err := net.Listen("tcp", config2.GRPC.HostPort)
	if err != nil {
		log.Panic(err)
	}

	mux := server.NewHTTPServer(config2.GRPC.Port, log)

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() (err error) {
		return grpcServer.Serve(listen)
	})
	g.Go(func() (err error) {
		return http.ListenAndServe(config2.HTTP.HostPort, mux)
	})

	err = g.Wait()
	if err != nil {
		log.Panic(err)
	}
}

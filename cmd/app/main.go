package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	config "grpcAvito/internal/config"
	"grpcAvito/internal/repository/postgres"
	"grpcAvito/internal/service"
	"grpcAvito/internal/usecase"
	"grpcAvito/internal/usecase/transactions"
	"grpcAvito/pkg/logging"
	"grpcAvito/pkg/server/grpc"
	httpServer "grpcAvito/pkg/server/http"
	"net"
	"net/http"
)

func init() {
	config.InitAll([]config.Config{
		postgres.Conf{},
		httpServer.Conf{},
		grpc.Conf{},
	})
}

func main() {
	log := logging.GetLogger()
	postgresDB, err := postgres.NewPostgresDB(postgres.Config)
	if err != nil {
		log.Fatal(fmt.Errorf("main - postgres.NewPostgresDB: %w", err))
	}

	postgresRepository := postgres.New(postgresDB, log)

	txService := transactions.New(postgresDB, log)

	useCases := usecase.New(postgresRepository, txService, log)

	services := service.New(useCases, log)

	grpcServer := grpc.New(services, log)

	listen, err := net.Listen("tcp", grpc.Config.HostPort)
	if err != nil {
		log.Fatal(fmt.Errorf("main - net.Listen: %w", err))
	}

	mux := httpServer.New(grpc.Config.Port, log)

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() (err error) {
		return grpcServer.Serve(listen)
	})
	g.Go(func() (err error) {
		serv := http.Server{
			Addr:         httpServer.Config.HostPort,
			ReadTimeout:  httpServer.Config.Duration,
			WriteTimeout: httpServer.Config.Duration,
			Handler:      mux,
		}
		return serv.ListenAndServe()
	})

	err = g.Wait()
	if err != nil {
		log.Panic(err)
	}
}

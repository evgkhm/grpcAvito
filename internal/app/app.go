package app

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"grpcAvito/internal/config"
	"grpcAvito/internal/repository/postgres"
	"grpcAvito/internal/service"
	"grpcAvito/internal/usecase"
	"grpcAvito/pkg/logging"
	"grpcAvito/pkg/server/grpc"
	httpServer "grpcAvito/pkg/server/http"
	"net"
	"net/http"
)

func Run() {
	log := logging.GetLogger()
	postgresDB, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.NewPostgresDB: %w", err))
	}

	postgresRepository := postgres.New(postgresDB, log)

	useCases := usecase.New(postgresRepository, log, postgresDB)

	services := service.New(useCases, log)

	grpcServer := grpc.New(services, log)

	listen, err := net.Listen("tcp", config.GRPC.HostPort)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - net.Listen: %w", err))
	}

	mux := httpServer.New(config.GRPC.Port, log)

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() (err error) {
		return grpcServer.Serve(listen)
	})
	g.Go(func() (err error) {
		serv := http.Server{
			Addr:         config.HTTP.HostPort,
			ReadTimeout:  config.HTTP.Duration,
			WriteTimeout: config.HTTP.Duration,
			Handler:      mux,
		}
		return serv.ListenAndServe()
	})

	err = g.Wait()
	if err != nil {
		log.Panic(err)
	}
}

package main

import (
	"context"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"grpcAvito/server/internal/config"
	postgresRepo "grpcAvito/server/internal/repository/postgres"
	"grpcAvito/server/internal/server"
	"grpcAvito/server/internal/service"
	"grpcAvito/server/internal/usecase"
	"grpcAvito/server/pkg/logging"
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
	postgresDB, err := postgresRepo.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to initialize database connection: %s", err.Error())
	}

	migrateDB(postgresDB, log)

	postgresRepository := postgresRepo.NewRepositoryPostgres(postgresDB, log)

	useCases := usecase.NewUseCase(postgresRepository, log)

	service := service.NewService(useCases, log)

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

func migrateDB(db *sqlx.DB, log *logrus.Logger) {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatalf("Couldn't get database instance for running migrations. %s", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("./db/migrations", "grpcAvito", driver)
	if err != nil {
		log.Fatalf("Couldn't create migrate instance. %s", err.Error())
	}

	if err := m.Up(); err != nil {
		log.Printf("Couldn't run database migration. %s", err.Error())
	} else {
		log.Println("Database migration was run successfully")
	}
}

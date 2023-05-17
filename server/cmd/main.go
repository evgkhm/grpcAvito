package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"grpcAvito/proto"
	"grpcAvito/server/internal/config"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	config.InitAll([]config.Config{
		config.PostgresConfig{},
		config.HttpConfig{},
	})
}

var (
	port = flag.Int("port", 50051, "The server port")
)

type Server struct {
	proto.UnimplementedServerServer
}

func (s *Server) Create(ctx context.Context, in *proto.CreateReq) (*proto.CreateReply, error) {
	return &proto.CreateReply{Id: "Hello again "}, nil //тут вывод айди и баланса
}

func main() {
	log := logging.GetLogger()
	postgresDB, err := postgresRepo.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to initialize database connection: %s", err.Error())
	}

	migrateDB(postgresDB, log)

	//flag.Parse()

	postgresRepository := postgresRepo.NewRepositoryPostgres(postgresDB, log)
	services := service.NewService(postgresRepository, postgresDB, log)
	handlers := handler.NewHandler(services, demoGateway, log)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterServerServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan

	log.Printf("Recieved terminate, %v\n", sig)

	ctx := context.Background()
	ctx, cancelCtx := context.WithTimeout(ctx, 30*time.Second)
	defer cancelCtx()

	if err := srv.Shutdown(ctx, log); err != nil {
		panic(err)
	}
}

func migrateDB(db *sqlx.DB, log *logrus.Logger) {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatalf("Couldn't get database instance for running migrations. %s", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "decide-database", driver)
	if err != nil {
		log.Fatalf("Couldn't create migrate instance. %s", err.Error())
	}

	if err := m.Up(); err != nil {
		log.Printf("Couldn't run database migration. %s", err.Error())
	} else {
		log.Println("Database migration was run successfully")
	}
}

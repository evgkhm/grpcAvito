package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpcAvito/proto"
	"log"
	"net"
)

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
	flag.Parse()
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
}

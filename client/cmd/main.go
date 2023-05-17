package main

import (
	"context"
	"google.golang.org/grpc"
	"grpcAvito/proto"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could't connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.Create(ctx, &proto.CreateReq{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetId())
}

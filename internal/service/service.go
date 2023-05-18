package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/usecase"
	"grpcAvito/proto"
)

type ServerServer interface {
	Create(context.Context, *proto.CreateReq) (*proto.CreateReply, error)
}

type Service struct {
	proto.UnimplementedServerServer
	useCase *usecase.UseCase
	log     *logrus.Logger
}

func (s Service) Create(ctx context.Context, req *proto.CreateReq) (*proto.CreateReply, error) {
	//TODO implement me
	s.log.Printf("Received: %v", req.GetUsername())
	return &proto.CreateReply{Id: "1 "}, nil
	//panic("implement me")
}

func NewService(useCase *usecase.UseCase, log *logrus.Logger) *Service {

	//txService := NewTransactionService(postgresDB, log)

	//usersService := NewUserService(repo, txService, log)
	return &Service{
		UnimplementedServerServer: proto.UnimplementedServerServer{},
		useCase:                   useCase,
		log:                       log,
	}
}

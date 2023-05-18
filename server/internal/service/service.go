package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"grpcAvito/proto"
	"grpcAvito/server/internal/entity"
	"grpcAvito/server/internal/usecase"
)

type ServerServer interface {
	Create(context.Context, *proto.CreateReq) (*proto.CreateReply, error)
	mustEmbedUnimplementedServerServer()
}

type Service struct {
	proto.UnimplementedServerServer
	useCase *usecase.UseCase
	log     *logrus.Logger
}

func (s Service) Create(ctx context.Context, req *proto.CreateReq) (*proto.CreateReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) mustEmbedUnimplementedServerServer() {
	//TODO implement me
	panic("implement me")
}

type UsersService interface {
	Create(user entity.User) (int64, error)
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

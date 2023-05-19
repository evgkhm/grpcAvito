package service

import (
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/usecase"
	"grpcAvito/proto"
)

type Service struct {
	proto.UnimplementedServerServer
	useCase *usecase.UseCase
	log     *logrus.Logger
}

func NewService(useCase *usecase.UseCase, log *logrus.Logger) *Service {
	return &Service{
		UnimplementedServerServer: proto.UnimplementedServerServer{},
		useCase:                   useCase,
		log:                       log,
	}
}
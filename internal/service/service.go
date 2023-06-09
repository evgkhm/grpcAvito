package service

import (
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/service/spec"
	"grpcAvito/internal/usecase"
)

type Service struct {
	spec.UnimplementedWalletAppServiceServer
	useCase usecase.Usecase
	log     *logrus.Logger
}

func New(useCase *usecase.UseCase, log *logrus.Logger) *Service {
	return &Service{
		UnimplementedWalletAppServiceServer: spec.UnimplementedWalletAppServiceServer{},
		useCase:                             useCase,
		log:                                 log,
	}
}

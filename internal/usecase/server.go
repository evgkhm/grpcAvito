package usecase

import (
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/repository/postgres"
)

type serverUseCase struct {
	repo *postgres.RepositoriesPostgres
	log  *logrus.Logger
}

func NewServerUseCase(repo *postgres.RepositoriesPostgres, log *logrus.Logger) *serverUseCase {
	return &serverUseCase{
		repo: repo,
		log:  log,
	}
}

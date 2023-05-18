package service

import (
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/repository/postgres"
	"grpcAvito/internal/usecase"
)

type UsersServiceImpl struct {
	repo      postgres.UsersRepository
	txService *usecase.TransactionServiceImpl
	log       *logrus.Logger
}

func NewUserService(repo postgres.UsersRepository, txService *usecase.TransactionServiceImpl, log *logrus.Logger) *UsersServiceImpl {
	return &UsersServiceImpl{
		repo:      repo,
		txService: txService,
		log:       log,
	}
}

/*func (s *UsersServiceImpl) CreateUser(user entity.User) (int64, error) {
	return 0, nil
}*/

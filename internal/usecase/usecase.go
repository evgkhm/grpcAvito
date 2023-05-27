package usecase

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/repository/postgres"
)

type UseCase struct {
	repo      *postgres.RepositoriesPostgres
	txService *TransactionServiceImpl
	log       *logrus.Logger
}

func New(repo *postgres.RepositoriesPostgres, log *logrus.Logger, postgresDB *sqlx.DB) *UseCase {
	txService := NewTransactionService(postgresDB, log)
	return &UseCase{
		repo:      repo,
		log:       log,
		txService: txService,
	}
}

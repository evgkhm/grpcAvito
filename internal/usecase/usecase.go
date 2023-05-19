package usecase

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/repository/postgres"
)

type ServerUseCase interface {
	Create(ctx context.Context, Id, Balance string) error
}

type UseCase struct {
	repo      *postgres.RepositoriesPostgres
	txService *TransactionServiceImpl
	log       *logrus.Logger
}

func (u UseCase) Create(ctx context.Context, Id uint32, Balance float32) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}
	defer u.txService.Commit(tx)

	//TODO: проверка что есть такой юзер

	err = u.repo.Create(ctx, tx, Id, Balance)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	return nil
}

func NewUseCase(repo *postgres.RepositoriesPostgres, log *logrus.Logger, postgresDB *sqlx.DB) *UseCase {
	txService := NewTransactionService(postgresDB, log)
	return &UseCase{
		repo:      repo,
		log:       log,
		txService: txService,
	}
}

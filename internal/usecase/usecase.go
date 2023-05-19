package usecase

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/entity"
	"grpcAvito/internal/repository/postgres"
)

type ServerUseCase interface {
	Create(ctx context.Context, Id, Balance string) error
	Sum(ctx context.Context, userDTO entity.User) error
}

type UseCase struct {
	repo      *postgres.RepositoriesPostgres
	txService *TransactionServiceImpl
	log       *logrus.Logger
}

func (u UseCase) Sum(ctx context.Context, userDTO entity.User) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}
	defer u.txService.Commit(tx)

	if userDTO.Balance < 0 {
		u.txService.Rollback(tx)
		errUserNegativeBalance := errors.New("you cannot add a negative balance")
		return errUserNegativeBalance
	}

	err = u.repo.Sum(ctx, tx, userDTO)
	if err != nil {
		u.txService.Rollback(tx)
		return err
	}

	return nil
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

package usecase

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/entity"
	"grpcAvito/internal/repository/postgres"
)

type ServerUseCase interface {
	Create(ctx context.Context, userDTO entity.User) error
	Sum(ctx context.Context, userDTO entity.User) error
	Reservation(ctx context.Context, reservation entity.UserReservation) error
	Dereservation(dereservation entity.UserReservation, tx *sqlx.Tx) error
	Report(ctx context.Context, year uint32, month uint32) error
}

type UseCase struct {
	repo      *postgres.RepositoriesPostgres
	txService *TransactionServiceImpl
	log       *logrus.Logger
}

func NewUseCase(repo *postgres.RepositoriesPostgres, log *logrus.Logger, postgresDB *sqlx.DB) *UseCase {
	txService := NewTransactionService(postgresDB, log)
	return &UseCase{
		repo:      repo,
		log:       log,
		txService: txService,
	}
}

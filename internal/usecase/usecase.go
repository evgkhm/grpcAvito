package usecase

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/entity"
	"grpcAvito/internal/repository/postgres"
)

type UseCase struct {
	repo      postgres.Repository
	txService *TransactionServiceImpl
	log       *logrus.Logger
}

type Usecase interface {
	Sum(ctx context.Context, userDTO *entity.User) error
	Create(ctx context.Context, userDTO *entity.User) error
	GetBalance(ctx context.Context, dto *entity.User) error
	Revenue(ctx context.Context, revenue *entity.UserRevenue) error
	DeleteReservation(ctx context.Context, reservation *entity.UserReservation) error
	Reservation(ctx context.Context, reservation *entity.UserReservation) error
	Report(ctx context.Context, year uint32, month uint32) error
}

func New(repo *postgres.Repo, log *logrus.Logger, postgresDB *sqlx.DB) *UseCase {
	txService := NewTransactionService(postgresDB, log)
	return &UseCase{
		repo:      repo,
		log:       log,
		txService: txService,
	}
}

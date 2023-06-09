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
	UserBalanceAccrual(ctx context.Context, userDTO *entity.User) error
	CreateUser(ctx context.Context, userDTO *entity.User) error
	GetBalance(ctx context.Context, dto *entity.User) error
	UserOrderRevenue(ctx context.Context, revenue *entity.UserRevenue) error
	UserOrderDeleteReservation(ctx context.Context, reservation *entity.UserReservation) error
	UserOrderReservation(ctx context.Context, reservation *entity.UserReservation) error
	CreateMonthReport(ctx context.Context, year uint32, month uint32) error
}

func New(repo *postgres.Repo, log *logrus.Logger, postgresDB *sqlx.DB) *UseCase {
	txService := NewTransactionService(postgresDB, log)
	return &UseCase{
		repo:      repo,
		log:       log,
		txService: txService,
	}
}

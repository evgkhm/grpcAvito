package usecase

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/entity"
	"grpcAvito/internal/repository/postgres"
)

type UseCase struct {
	User
	Order
	Report
}

type User interface {
	CreateUser(ctx context.Context, userDTO *entity.User) error
	GetBalance(ctx context.Context, dto *entity.User) error
	UserBalanceAccrual(ctx context.Context, userDTO *entity.User) error
}

type Order interface {
	UserOrderReservation(ctx context.Context, reservation *entity.UserReservation) error
	UserOrderDeleteReservation(ctx context.Context, reservation *entity.UserReservation) error
	UserOrderRevenue(ctx context.Context, revenue *entity.UserRevenue) error
}

type Report interface {
	CreateMonthReport(ctx context.Context, year uint32, month uint32) error
}

func New(repo *postgres.Repository, log *logrus.Logger, postgresDB *sqlx.DB) *UseCase {
	txService := NewTransactionService(postgresDB, log)
	return &UseCase{
		User:   NewUserUseCase(repo.UserRepository, repo.OrderRepository, repo.ReportRepository, log, txService),
		Order:  NewOrderUseCase(repo.UserRepository, repo.OrderRepository, repo.ReportRepository, log, txService),
		Report: NewReportUseCase(repo.UserRepository, repo.OrderRepository, repo.ReportRepository, log, txService),
	}
}

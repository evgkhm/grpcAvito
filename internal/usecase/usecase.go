package usecase

import (
	"context"
	"github.com/sirupsen/logrus"
	user2 "grpcAvito/internal/entity/user"
	"grpcAvito/internal/repository/postgres"
	"grpcAvito/internal/usecase/order"
	"grpcAvito/internal/usecase/report"
	"grpcAvito/internal/usecase/transactions"
	"grpcAvito/internal/usecase/user"
)

type UseCase struct {
	User
	Order
	Report
}

type User interface {
	CreateUser(ctx context.Context, userDTO *user2.User) error
	GetBalance(ctx context.Context, dto *user2.User) error
	UserBalanceAccrual(ctx context.Context, userDTO *user2.User) error
}

type Order interface {
	UserOrderReservation(ctx context.Context, reservation *user2.Reservation) error
	UserOrderDeleteReservation(ctx context.Context, reservation *user2.Reservation) error
	UserOrderRevenue(ctx context.Context, revenue *user2.Revenue) error
}

type Report interface {
	CreateMonthReport(ctx context.Context, year uint32, month uint32) error
}

func New(repo *postgres.Repository, txService *transactions.TransactionServiceImpl, log *logrus.Logger) *UseCase {
	return &UseCase{
		User:   user.NewUserUseCase(repo.UserRepository, repo.OrderRepository, repo.ReportRepository, log, txService),
		Order:  order.NewOrderUseCase(repo.UserRepository, repo.OrderRepository, repo.ReportRepository, log, txService),
		Report: report.NewReportUseCase(repo.UserRepository, repo.OrderRepository, repo.ReportRepository, log, txService),
	}
}

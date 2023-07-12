package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	user2 "grpcAvito/internal/entity/user"
	"grpcAvito/internal/repository/postgres/order"
	"grpcAvito/internal/repository/postgres/report"
	"grpcAvito/internal/repository/postgres/user"
)

type Repository struct {
	UserRepository
	OrderRepository
	ReportRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, tx *sqlx.Tx, userDTO *user2.User) error
	GetBalance(ctx context.Context, tx *sqlx.Tx, userDTO *user2.User) (float32, error)
	UserBalanceAccrual(ctx context.Context, tx *sqlx.Tx, userDTO *user2.User) error
	MinusBalance(ctx context.Context, tx *sqlx.Tx, userDTO *user2.User) error
}

type OrderRepository interface {
	UserOrderReservation(ctx context.Context, tx *sqlx.Tx, reservation *user2.Reservation) error
	UserOrderRevenue(ctx context.Context, tx *sqlx.Tx, revenue *user2.Revenue) error
	UserOrderDeleteReservation(ctx context.Context, tx *sqlx.Tx, reservation *user2.Reservation) error
}

type ReportRepository interface {
	CreateMonthReport(ctx context.Context, tx *sqlx.Tx, year uint32, month uint32) (map[uint32]float32, error)
}

func New(db *sqlx.DB, log *logrus.Logger) *Repository {
	return &Repository{
		UserRepository:   user.NewUserRepository(db, log),
		OrderRepository:  order.NewOrderRepository(db, log),
		ReportRepository: report.NewReportRepository(db, log),
	}
}

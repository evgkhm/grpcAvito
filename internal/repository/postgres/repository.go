package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/entity"
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
	CreateUser(ctx context.Context, tx *sqlx.Tx, userDTO *entity.User) error
	GetBalance(ctx context.Context, tx *sqlx.Tx, userDTO *entity.User) (float32, error)
	UserBalanceAccrual(ctx context.Context, tx *sqlx.Tx, userDTO *entity.User) error
	MinusBalance(ctx context.Context, tx *sqlx.Tx, userDTO *entity.User) error
}

type OrderRepository interface {
	UserOrderReservation(ctx context.Context, tx *sqlx.Tx, reservation *entity.UserReservation) error
	UserOrderRevenue(ctx context.Context, tx *sqlx.Tx, revenue *entity.UserRevenue) error
	UserOrderDeleteReservation(ctx context.Context, tx *sqlx.Tx, reservation *entity.UserReservation) error
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

package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/entity"
)

type Repository struct {
	UserRepository
	OrderRepository
	ReportRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, tx *sqlx.Tx, userDTO *entity.User) error
	UserBalanceAccrual(ctx context.Context, tx *sqlx.Tx, userDTO *entity.User) error
	GetBalance(ctx context.Context, tx *sqlx.Tx, userDTO *entity.User) (float32, error)
}

type OrderRepository interface {
	UserOrderRevenue(ctx context.Context, tx *sqlx.Tx, revenue *entity.UserRevenue) error
	UserOrderReservation(ctx context.Context, tx *sqlx.Tx, reservation *entity.UserReservation) error
	MinusBalance(ctx context.Context, tx *sqlx.Tx, userDTO *entity.User) error
	UserOrderDeleteReservation(ctx context.Context, tx *sqlx.Tx, reservation *entity.UserReservation) error
}

type ReportRepository interface {
	CreateMonthReport(ctx context.Context, tx *sqlx.Tx, year uint32, month uint32) (map[uint32]float32, error)
}

func New(db *sqlx.DB, log *logrus.Logger) *Repository {
	return &Repository{
		UserRepository:   NewUserRepository(db, log),
		OrderRepository:  NewOrderRepository(db, log),
		ReportRepository: NewReportRepository(db, log),
	}
}

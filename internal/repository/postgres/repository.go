package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/entity"
)

type UsersRepository interface {
	Create(ctx context.Context, tx *sqlx.Tx, userDTO *entity.User) error
	Sum(ctx context.Context, tx *sqlx.Tx, userDTO *entity.User) error
	GetBalance(ctx context.Context, tx *sqlx.Tx, userDTO *entity.User) (float32, error)
}

type ReservationRepository interface {
	Reservation(ctx context.Context, tx *sqlx.Tx, reservation *entity.UserReservation) error
	MinusBalance(ctx context.Context, tx *sqlx.Tx, userDTO *entity.User) error
	DeleteReservation(ctx context.Context, tx *sqlx.Tx, reservation *entity.UserReservation) error
}

type RevenueRepository interface {
	Revenue(ctx context.Context, tx *sqlx.Tx, revenue *entity.UserRevenue) error
}

type ReportRepository interface {
	GetReport(ctx context.Context, tx *sqlx.Tx, year uint32, month uint32) (map[uint32]float32, error)
}

type RepositoriesPostgres struct {
	UsersRepository
	ReservationRepository
	RevenueRepository
	ReportRepository
}

func New(db *sqlx.DB, log *logrus.Logger) *RepositoriesPostgres {

	return &RepositoriesPostgres{
		UsersRepository:       NewUsersPostgres(db, log),
		ReservationRepository: NewReservationRepository(db, log),
		RevenueRepository:     NewRevenueRepository(db, log),
		ReportRepository:      NewReportRepository(db, log),
	}
}

package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	entity "grpcAvito/internal/entity"
)

type UsersRepository interface {
	Create(ctx context.Context, tx *sqlx.Tx, Id uint32, Balance float32) error
	Sum(ctx context.Context, tx *sqlx.Tx, userDTO entity.User) error
}

type ReservationRepository interface {
	Reservation(reservation entity.UserReservation, tx *sqlx.Tx) error
}

type RevenueRepository interface {
	Revenue(revenue entity.UserRevenue, tx *sqlx.Tx) error
}

type RepositoriesPostgres struct {
	UsersRepository
	ReservationRepository
	RevenueRepository
}

func NewRepositoryPostgres(db *sqlx.DB, log *logrus.Logger) *RepositoriesPostgres {

	return &RepositoriesPostgres{
		UsersRepository:       NewUsersPostgres(db, log),
		ReservationRepository: NewReservationRepository(db, log),
		RevenueRepository:     NewRevenueRepository(db, log),
	}
}

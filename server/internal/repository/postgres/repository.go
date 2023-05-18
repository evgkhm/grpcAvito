package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"grpcAvito/server/internal/entity"
)

type UsersRepository interface {
	CreateUser(user entity.User, tx *sqlx.Tx) (int64, error)
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

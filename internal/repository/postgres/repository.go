package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	entity2 "grpcAvito/internal/entity"
)

type UsersRepository interface {
	CreateUser(user entity2.User, tx *sqlx.Tx) (int64, error)
}

type ReservationRepository interface {
	Reservation(reservation entity2.UserReservation, tx *sqlx.Tx) error
}

type RevenueRepository interface {
	Revenue(revenue entity2.UserRevenue, tx *sqlx.Tx) error
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

package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/entity"
)

type ReservationRepositoryImpl struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewReservationRepository(db *sqlx.DB, log *logrus.Logger) *ReservationRepositoryImpl {
	return &ReservationRepositoryImpl{
		db:  db,
		log: log,
	}
}

func (r ReservationRepositoryImpl) Reservation(reservation entity.UserReservation, tx *sqlx.Tx) error {
	return nil
}

package postgres

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
	var idOrder uint32
	query := fmt.Sprintf(
		"INSERT INTO %s (id, balance) "+
			"VALUES ($1, $2) RETURNING id", reservationTable)

	row := tx.QueryRow(query, reservation.Id, reservation.IdService, reservation.IdOrder, reservation.Cost)
	if err, ok := row.Scan(&idOrder).(*pq.Error); ok {
		if err.Code == "23505" {
			ErrUserAlreadyExist := errors.New("such idOrder already exists")
			return ErrUserAlreadyExist
		}
		return err
	}

	return nil
}

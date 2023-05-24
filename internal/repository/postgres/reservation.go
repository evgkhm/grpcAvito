package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
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

func (r ReservationRepositoryImpl) Reservation(reservation *entity.UserReservation, tx *sqlx.Tx) error {
	var idOrder uint32
	query := fmt.Sprintf(
		"INSERT INTO %s (id, id_service, id_order, cost) "+
			"VALUES ($1, $2, $3, $4) RETURNING id", reservationTable)

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

func (r ReservationRepositoryImpl) MinusBalance(tx *sqlx.Tx, user *entity.User) error {
	query := `UPDATE usr SET "balance"=$1 WHERE "id"=$2`
	_, err := tx.Exec(query, user.Balance, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r ReservationRepositoryImpl) Dereservation(dereservation entity.UserReservation, tx *sqlx.Tx) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 and id_service=$2 and id_order=$3 and cost=$4", reservationTable)
	_, err := tx.Exec(query, dereservation.Id, dereservation.IdService, dereservation.IdOrder, dereservation.Cost)
	if err != nil {
		return err
	}

	return nil
}

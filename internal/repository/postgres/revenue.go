package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/entity"
)

type RevenueRepositoryImpl struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewRevenueRepository(db *sqlx.DB, log *logrus.Logger) *RevenueRepositoryImpl {
	return &RevenueRepositoryImpl{
		db:  db,
		log: log,
	}
}

func (r RevenueRepositoryImpl) MinusReservation(reservation entity.UserReservation, tx *sqlx.Tx) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 and id_service=$2 and id_order=$3 and cost=$4", reservationTable)
	_, err := tx.Exec(query, reservation.Id, reservation.IdService, reservation.IdOrder, reservation.Cost)
	if err != nil {
		return err
	}
	return nil
}

func (r RevenueRepositoryImpl) Revenue(revenue entity.UserRevenue, tx *sqlx.Tx) error {
	var idOrder uint32
	query := fmt.Sprintf(
		"INSERT INTO %s (id, balance) "+
			"VALUES ($1, $2) RETURNING id", revenueTable)

	row := tx.QueryRow(query, revenue.Id, revenue.IdService, revenue.IdOrder, revenue.Cost)
	if err, ok := row.Scan(&idOrder).(*pq.Error); ok {
		if err.Code == "23505" {
			ErrUserAlreadyExist := errors.New("such idOrder already exists")
			return ErrUserAlreadyExist
		}
		return err
	}

	return nil
}

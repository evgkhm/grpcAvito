package postgres

import (
	"context"
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

func (r ReservationRepositoryImpl) Reservation(ctx context.Context, tx *sqlx.Tx, reservation *entity.UserReservation) error {
	var idOrder uint32

	query := `INSERT INTO reservation 
    	(id, id_service, id_order, cost) 
			VALUES ($1, $2, $3, $4) RETURNING id`

	row := tx.QueryRowxContext(ctx, query, reservation.Id, reservation.IdService, reservation.IdOrder, reservation.Cost)
	if err, ok := row.Scan(&idOrder).(*pq.Error); ok {
		if err.Code == "23505" {
			return fmt.Errorf("postgres: Reservation: QueryRowxContext: Scan: %w", errUserNotExist)
		}
		return fmt.Errorf("postgres: Reservation: QueryRowxContext: Scan: %w", err)
	}

	return nil
}

func (r ReservationRepositoryImpl) MinusBalance(ctx context.Context, tx *sqlx.Tx, user *entity.User) error {
	query := `UPDATE usr SET "balance"=$1 WHERE "id"=$2`
	_, err := tx.ExecContext(ctx, query, user.Balance, user.Id)
	if err != nil {
		return fmt.Errorf("postgres: MinusBalance: ExecContext: %w", err)
	}
	return nil
}

func (r ReservationRepositoryImpl) DeleteReservation(ctx context.Context, tx *sqlx.Tx, reservation *entity.UserReservation) error {
	query := `DELETE FROM reservation 
       WHERE id=$1 and id_service=$2 and id_order=$3 and cost=$4`
	_, err := tx.ExecContext(ctx, query, reservation.Id, reservation.IdService, reservation.IdOrder, reservation.Cost)
	if err != nil {
		return fmt.Errorf("postgres: DeleteReservation: ExecContext: %w", err)
	}

	return nil
}

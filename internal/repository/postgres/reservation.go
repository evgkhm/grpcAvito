package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"grpcAvito/internal/entity"
)

func (r Repo) Reservation(ctx context.Context, tx *sqlx.Tx, reservation *entity.UserReservation) error {
	var idOrder uint32

	query := `INSERT INTO reservation 
    	(id, id_service, id_order, cost) 
			VALUES ($1, $2, $3, $4) RETURNING id`

	row := tx.QueryRowxContext(ctx, query, reservation.Id, reservation.IdService, reservation.IdOrder, reservation.Cost)
	if err, ok := row.Scan(&idOrder).(*pq.Error); ok {
		if err.Code == "23505" {
			return fmt.Errorf("postgres - ReservationRepositoryImpl - Reservation - tx.QueryRowxContext - row.Scan: %w", errUserNotExist)
		}
		return fmt.Errorf("postgres - ReservationRepositoryImpl - Reservation - tx.QueryRowxContext - row.Scan: %w", err)
	}

	return nil
}

func (r Repo) MinusBalance(ctx context.Context, tx *sqlx.Tx, user *entity.User) error {
	query := `UPDATE usr SET "balance"=$1 WHERE "id"=$2`
	_, err := tx.ExecContext(ctx, query, user.Balance, user.Id)
	if err != nil {
		return fmt.Errorf("postgres - ReservationRepositoryImpl - MinusBalance - tx.ExecContext: %w", err)
	}
	return nil
}

func (r Repo) DeleteReservation(ctx context.Context, tx *sqlx.Tx, reservation *entity.UserReservation) error {
	query := `DELETE FROM reservation 
       WHERE id=$1 and id_service=$2 and id_order=$3 and cost=$4`
	_, err := tx.ExecContext(ctx, query, reservation.Id, reservation.IdService, reservation.IdOrder, reservation.Cost)
	if err != nil {
		return fmt.Errorf("postgres - ReservationRepositoryImpl - DeleteReservation - tx.ExecContext: %w", err)
	}

	return nil
}

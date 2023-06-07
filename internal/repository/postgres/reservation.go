package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"grpcAvito/internal/entity"
)

func (r Repo) UserOrderReservation(ctx context.Context, tx *sqlx.Tx, reservation *entity.UserReservation) error {
	var idOrder uint32
	var duplicateEntryError = &pq.Error{Code: "23505"}
	query := `INSERT INTO reservation 
    	(id, id_service, id_order, cost) 
			VALUES ($1, $2, $3, $4) RETURNING id`

	row := tx.QueryRowxContext(ctx, query, reservation.ID, reservation.IDService, reservation.IDOrder, reservation.Cost)
	if err, ok := row.Scan(&idOrder).(*pq.Error); ok {
		if errors.As(err, &duplicateEntryError) {
			return fmt.Errorf("postgres - ReservationRepositoryImpl - UserOrderReservation - tx.QueryRowxContext - row.Scan: %w", ErrOrderAlreadyExist)
		}
		return fmt.Errorf("postgres - ReservationRepositoryImpl - UserOrderReservation - tx.QueryRowxContext - row.Scan: %w", err)
	}

	return nil
}

func (r Repo) MinusBalance(ctx context.Context, tx *sqlx.Tx, user *entity.User) error {
	query := `UPDATE usr SET "balance"=$1 WHERE "id"=$2`
	_, err := tx.ExecContext(ctx, query, user.Balance, user.ID)
	if err != nil {
		return fmt.Errorf("postgres - ReservationRepositoryImpl - MinusBalance - tx.ExecContext: %w", err)
	}
	return nil
}

func (r Repo) UserOrderDeleteReservation(ctx context.Context, tx *sqlx.Tx, reservation *entity.UserReservation) error {
	query := `DELETE FROM reservation 
       WHERE id=$1 and id_service=$2 and id_order=$3 and cost=$4`
	_, err := tx.ExecContext(ctx, query, reservation.ID, reservation.IDService, reservation.IDOrder, reservation.Cost)
	if err != nil {
		return fmt.Errorf("postgres - ReservationRepositoryImpl - UserOrderDeleteReservation - tx.ExecContext: %w", err)
	}

	return nil
}

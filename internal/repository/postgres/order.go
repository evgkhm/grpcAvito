package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/entity"
)

type OrderRepo struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewOrderRepository(db *sqlx.DB, log *logrus.Logger) *OrderRepo {
	return &OrderRepo{
		db:  db,
		log: log,
	}
}

func (o OrderRepo) UserOrderReservation(ctx context.Context, tx *sqlx.Tx, reservation *entity.UserReservation) error {
	var idOrder uint32
	var duplicateEntryError = &pq.Error{Code: "23505"}
	query := `INSERT INTO reservation 
    	(id, service_id, order_id, cost) 
			VALUES ($1, $2, $3, $4) RETURNING id`

	row := tx.QueryRowxContext(ctx, query, reservation.ID, reservation.ServiceID, reservation.OrderID, reservation.Cost)
	err := row.Scan(&idOrder)
	switch {
	case errors.As(err, &duplicateEntryError):
		return ErrOrderAlreadyExist // return fmt.Errorf("postgres - ReservationRepositoryImpl - UserOrderReservation - tx.QueryRowxContext - row.Scan: %w", ErrOrderAlreadyExist)
	case err != nil:
		return fmt.Errorf("postgres - ReservationRepositoryImpl - UserOrderReservation - tx.QueryRowxContext - row.Scan: %w", err)
	}
	return nil
}

func (o OrderRepo) MinusBalance(ctx context.Context, tx *sqlx.Tx, user *entity.User) error {
	query := `UPDATE usr SET "balance"=$1 WHERE "id"=$2`
	_, err := tx.ExecContext(ctx, query, user.Balance, user.ID)
	if err != nil {
		return fmt.Errorf("postgres - ReservationRepositoryImpl - MinusBalance - tx.ExecContext: %w", err)
	}
	return nil
}

func (o OrderRepo) UserOrderDeleteReservation(ctx context.Context, tx *sqlx.Tx, reservation *entity.UserReservation) error {
	query := `DELETE FROM reservation 
       WHERE id=$1 and service_id=$2 and order_id=$3 and cost=$4`
	_, err := tx.ExecContext(ctx, query, reservation.ID, reservation.ServiceID, reservation.OrderID, reservation.Cost)
	if err != nil {
		return fmt.Errorf("postgres - ReservationRepositoryImpl - UserOrderDeleteReservation - tx.ExecContext: %w", err)
	}
	return nil
}

func (o OrderRepo) UserOrderRevenue(ctx context.Context, tx *sqlx.Tx, revenue *entity.UserRevenue) error {
	var idOrder uint32
	var duplicateEntryError = &pq.Error{Code: "23505"}
	query := `INSERT INTO revenue 
    	(id, service_id, order_id, cost)
			VALUES ($1, $2, $3, $4) RETURNING id`
	row := tx.QueryRowxContext(ctx, query, revenue.ID, revenue.ServiceID, revenue.OrderID, revenue.Cost)
	err := row.Scan(&idOrder)
	switch {
	case errors.As(err, &duplicateEntryError):
		return ErrRevenueAlreadyExist // return fmt.Errorf("postgres - RevenueRepositoryImpl - UserOrderRevenue - tx.QueryRowxContext - row.Scan: %w", ErrRevenueAlreadyExist)
	case err != nil:
		return fmt.Errorf("postgres - RevenueRepositoryImpl - UserOrderRevenue - tx.QueryRowxContext - row.Scan: %w", err)
	}
	return nil
}

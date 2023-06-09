package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"grpcAvito/internal/entity"
)

func (r Repo) UserOrderRevenue(ctx context.Context, tx *sqlx.Tx, revenue *entity.UserRevenue) error {
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

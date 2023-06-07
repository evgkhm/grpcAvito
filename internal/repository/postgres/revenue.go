package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"grpcAvito/internal/entity"
)

func (r Repo) Revenue(ctx context.Context, tx *sqlx.Tx, revenue *entity.UserRevenue) error {
	var idOrder uint32
	var duplicateEntryError = &pq.Error{Code: "23505"}
	query := `INSERT INTO revenue 
    	(id, id_service, id_order, cost)
			VALUES ($1, $2, $3, $4) RETURNING id`
	row := tx.QueryRowxContext(ctx, query, revenue.ID, revenue.IDService, revenue.IDOrder, revenue.Cost)
	if err, ok := row.Scan(&idOrder).(*pq.Error); ok {
		if errors.As(err, &duplicateEntryError) {
			return fmt.Errorf("postgres - RevenueRepositoryImpl - UserOrderRevenue - tx.QueryRowxContext - row.Scan: %w", ErrRevenueAlreadyExist)
		}
		return fmt.Errorf("postgres - RevenueRepositoryImpl - UserOrderRevenue - tx.QueryRowxContext - row.Scan: %w", err)
	}

	return nil
}

package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"grpcAvito/internal/entity"
)

func (r Repo) Revenue(ctx context.Context, tx *sqlx.Tx, revenue *entity.UserRevenue) error {
	var idOrder uint32
	query := `INSERT INTO revenue 
    	(id, id_service, id_order, cost)
			VALUES ($1, $2, $3, $4) RETURNING id`
	row := tx.QueryRowxContext(ctx, query, revenue.Id, revenue.IdService, revenue.IdOrder, revenue.Cost)
	if err, ok := row.Scan(&idOrder).(*pq.Error); ok {
		if err.Code == "23505" {
			return fmt.Errorf("postgres - RevenueRepositoryImpl - UserOrderRevenue - tx.QueryRowxContext - row.Scan: %w", errUserAlreadyExist)
		}
		return fmt.Errorf("postgres - RevenueRepositoryImpl - UserOrderRevenue - tx.QueryRowxContext - row.Scan: %w", err)
	}

	return nil
}

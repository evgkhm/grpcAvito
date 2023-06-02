package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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

func (r RevenueRepositoryImpl) Revenue(ctx context.Context, tx *sqlx.Tx, revenue *entity.UserRevenue) error {
	var idOrder uint32
	query := `INSERT INTO revenue 
    	(id, id_service, id_order, cost)
			VALUES ($1, $2, $3, $4) RETURNING id`
	row := tx.QueryRowxContext(ctx, query, revenue.Id, revenue.IdService, revenue.IdOrder, revenue.Cost)
	if err, ok := row.Scan(&idOrder).(*pq.Error); ok {
		if err.Code == "23505" {
			return fmt.Errorf("postgres: Revenue: QueryRowxContext: Scan: %w", errUserAlreadyExist)
		}
		return fmt.Errorf("postgres: Revenue: QueryRowxContext: %w", err)
	}

	return nil
}

package postgres

import (
	"context"
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

/*func (r RevenueRepositoryImpl) DeleteReservation(ctx context.Context, reservation entity.UserReservation, tx *sqlx.Tx) error {
	query := `DELETE FROM reservation
       WHERE id=$1 and id_service=$2 and id_order=$3 and cost=$4`
	_, err := tx.ExecContext(ctx, query, reservation.Id, reservation.IdService, reservation.IdOrder, reservation.Cost)
	if err != nil {
		return err
	}
	return nil
}*/

func (r RevenueRepositoryImpl) Revenue(ctx context.Context, revenue entity.UserRevenue, tx *sqlx.Tx) error {
	var idOrder uint32
	query := `INSERT INTO revenue 
    	(id, id_service, id_order, cost)
			VALUES ($1, $2, $3, $4) RETURNING id`
	row := tx.QueryRowxContext(ctx, query, revenue.Id, revenue.IdService, revenue.IdOrder, revenue.Cost)
	if err, ok := row.Scan(&idOrder).(*pq.Error); ok {
		if err.Code == "23505" {
			return ErrUserAlreadyExist
		}
		return err
	}

	return nil
}

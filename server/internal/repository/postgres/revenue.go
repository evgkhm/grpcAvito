package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"grpcAvito/server/internal/entity"
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

func (r RevenueRepositoryImpl) Revenue(revenue entity.UserRevenue, tx *sqlx.Tx) error {
	return nil
}

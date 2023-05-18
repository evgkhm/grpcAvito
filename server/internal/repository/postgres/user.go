package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"grpcAvito/server/internal/entity"
)

type UsersRepositoryImpl struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewUsersPostgres(db *sqlx.DB, log *logrus.Logger) *UsersRepositoryImpl {
	return &UsersRepositoryImpl{
		db:  db,
		log: log,
	}
}

func (r UsersRepositoryImpl) CreateUser(user entity.User, tx *sqlx.Tx) (int64, error) {
	return 0, nil
}

package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
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

func (r UsersRepositoryImpl) Create(ctx context.Context, tx *sqlx.Tx, Id uint32, Balance float32) error {
	var id int64

	query := fmt.Sprintf(
		"INSERT INTO %s (id, balance) "+
			"VALUES ($1, $2) RETURNING id", usersTable)

	row := tx.QueryRow(query, Id, Balance)
	if err, ok := row.Scan(&id).(*pq.Error); ok {
		if err.Code == "23505" {
			ErrUserAlreadyExist := errors.New("такой пользователь уже существует")
			return ErrUserAlreadyExist
		}
		return err
	}
	return nil
}

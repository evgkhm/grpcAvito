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

func (r UsersRepositoryImpl) GetBalance(ctx context.Context, tx *sqlx.Tx, user entity.User) (float32, error) {
	var balance float32
	query := fmt.Sprintf("SELECT id FROM %s WHERE id=$1", usersTable)
	err := tx.Get(&balance, query, user.Id)
	return balance, err
}

func (r UsersRepositoryImpl) Create(ctx context.Context, tx *sqlx.Tx, user entity.User) error {
	var id int64

	query := fmt.Sprintf(
		"INSERT INTO %s (id, balance) "+
			"VALUES ($1, $2) RETURNING id", usersTable)

	row := tx.QueryRow(query, user.Id, user.Balance)
	if err, ok := row.Scan(&id).(*pq.Error); ok {
		if err.Code == "23505" {
			ErrUserAlreadyExist := errors.New("such user already exists")
			return ErrUserAlreadyExist
		}
		return err
	}
	return nil
}

func (r UsersRepositoryImpl) Sum(ctx context.Context, tx *sqlx.Tx, user entity.User) error {
	query := `UPDATE usr SET "balance"=$1 WHERE "id"=$2`
	_, err := tx.Exec(query, user.Balance, user.Id)
	if err != nil {
		return err
	}
	return nil
}

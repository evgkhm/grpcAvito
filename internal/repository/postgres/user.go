package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/entity"
)

var (
	ErrUserAlreadyExist = errors.New("such user already exists")
	ErrUserNotExist     = errors.New("such user does not exist")
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

func (r UsersRepositoryImpl) GetBalance(ctx context.Context, tx *sqlx.Tx, user *entity.User) (float32, error) {
	var balance float32
	query := `SELECT balance FROM usr WHERE id=$1 `
	err := tx.GetContext(ctx, &balance, query, user.Id)
	return balance, err
}

func (r UsersRepositoryImpl) Create(ctx context.Context, tx *sqlx.Tx, user entity.User) error {
	var id int64
	query := `INSERT INTO usr (id, balance) VALUES ($1, $2) RETURNING id`
	row := tx.QueryRowxContext(ctx, query, user.Id, user.Balance)
	if err, ok := row.Scan(&id).(*pq.Error); ok {
		if err.Code == "23505" {
			return ErrUserAlreadyExist
		}
		return err
	}
	return nil
}

func (r UsersRepositoryImpl) Sum(ctx context.Context, tx *sqlx.Tx, user *entity.User) error {
	query := `UPDATE usr SET "balance"=$1 WHERE "id"=$2`
	_, err := tx.ExecContext(ctx, query, user.Balance, user.Id)
	if err != nil {
		return err
	}
	return nil
}

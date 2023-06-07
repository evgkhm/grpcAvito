package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"grpcAvito/internal/entity"
)

func (r Repo) GetBalance(ctx context.Context, tx *sqlx.Tx, user *entity.User) (float32, error) {
	var balance float32
	query := `SELECT balance FROM usr WHERE id=$1 `
	err := tx.GetContext(ctx, &balance, query, user.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("postgres - UsersRepositoryImpl - GetBalance - tx.GetContext: %w", ErrUserNotExist)
		}
		return 0, fmt.Errorf("postgres - UsersRepositoryImpl - GetBalance - tx.GetContext: %w", err)
	}
	return balance, nil
}

func (r Repo) Create(ctx context.Context, tx *sqlx.Tx, user *entity.User) error {
	var id int64
	var duplicateEntryError = &pq.Error{Code: "23505"}
	query := `INSERT INTO usr (id, balance) VALUES ($1, $2) RETURNING id`
	row := tx.QueryRowxContext(ctx, query, user.ID, user.Balance)
	if err, ok := row.Scan(&id).(*pq.Error); ok {
		if errors.As(err, &duplicateEntryError) {
			return fmt.Errorf("postgres - UsersRepositoryImpl - Create - tx.QueryRowxContext - row.Scan: %w", ErrUserAlreadyExist)
		}
		return fmt.Errorf("postgres - UsersRepositoryImpl - Create - tx.QueryRowxContext - row.Scan: %w", err)
	}
	return nil
}

func (r Repo) Sum(ctx context.Context, tx *sqlx.Tx, user *entity.User) error {
	query := `UPDATE usr SET "balance"=$1 WHERE "id"=$2`
	_, err := tx.ExecContext(ctx, query, user.Balance, user.ID)
	if err != nil {
		return fmt.Errorf("postgres - UsersRepositoryImpl - Sum - tx.ExecContext: %w", err)
	}
	return nil
}

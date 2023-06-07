package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"grpcAvito/internal/config"
)

func NewPostgresDB() (*sqlx.DB, error) {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Name, config.Postgres.Pass, config.Postgres.SSLMode)

	db, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgresDB - sqlx.Open: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgresDB - db.Ping: %w", err)
	}

	return db, nil
}

package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"grpcAvito/internal/config"
)

const (
	usersTable       = "usr"
	reservationTable = "reservation"
	revenueTable     = "revenue"
)

func NewPostgresDB() (*sqlx.DB, error) {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Name, config.Postgres.Pass, config.Postgres.SSLMode)

	db, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

package transactions

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TransactionServiceImpl struct {
	postgresDB *sqlx.DB
	log        *logrus.Logger
}

func New(postgresDB *sqlx.DB, log *logrus.Logger) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		postgresDB: postgresDB,
		log:        log,
	}
}

func (t *TransactionServiceImpl) NewTransaction() (*sqlx.Tx, error) {
	return t.postgresDB.Beginx()
}

func (t *TransactionServiceImpl) Rollback(tx *sqlx.Tx) error {
	if err := tx.Rollback(); err != nil {
		return err
	}
	return nil
}

func (t TransactionServiceImpl) Commit(tx *sqlx.Tx) error {
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

package transactions

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TransactionServiceImpl struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewTransactionService(db *sqlx.DB, log *logrus.Logger) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		db:  db,
		log: log,
	}
}

func (u *TransactionServiceImpl) NewTransaction() (*sqlx.Tx, error) {
	return u.db.Beginx()
}

func (u *TransactionServiceImpl) Rollback(tx *sqlx.Tx) error {
	if err := tx.Rollback(); err != nil {
		return err
	}
	return nil
}

func (u TransactionServiceImpl) Commit(tx *sqlx.Tx) error {
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

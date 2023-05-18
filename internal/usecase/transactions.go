package usecase

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

func (s *TransactionServiceImpl) NewTransaction() (*sqlx.Tx, error) {
	return s.db.Beginx()
}

func (s *TransactionServiceImpl) Rollback(tx *sqlx.Tx) error {
	if err := tx.Rollback(); err != nil {
		s.log.Printf("service: unable to rollback transaction. %s", err.Error())
		return err
	}
	return nil
}

func (s TransactionServiceImpl) Commit(tx *sqlx.Tx) error {
	if err := tx.Commit(); err != nil {
		s.log.Printf("service: unable to commit transaction. %s", err.Error())
		return err
	}
	return nil
}

package repository

import (
	"database/sql"
	"time"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}
}

func (tdb *TransactionRepositoryDb) Insert(id string, accouintId string, amount float64, status string, errorMessage string) error {
	stmt, err := tdb.db.Prepare(`
		Insert into transactions (id, account_id, amount, status, error_message, created_at, updated_at)
		values($1, $2, $3, $4, $5, $6, $7)
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, accouintId, amount, status, errorMessage, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}
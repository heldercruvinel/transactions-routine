package pgtransactions

import (
	"database/sql"
	"log/slog"

	"github.com/heldercruvinel/transactions-routine/internal/transactions"
)

type DB struct {
	db *sql.DB
}

var _ transactions.DB = &DB{}

func GetDB(sqlDB *sql.DB) *DB {
	return &DB{
		db: sqlDB,
	}
}

// PostgreDB method to insert new transactions
func (pg *DB) Insert(transaction transactions.Transaction) (transactions.Transaction, error) {
	var result transactions.Transaction

	// Create new statement
	stmt, err := pg.db.Prepare(`
		INSERT INTO transactions.transactions (operation_id, account_id, amount, created_at) VALUES
		($1, $2, $3, now())
		RETURNING id, operation_id, account_id, amount, created_at;
	`)
	if err != nil {
		slog.Error("error to create transaction insert statement", slog.String("error", err.Error()))
		return result, err
	}
	defer stmt.Close()

	// Execute the statement, if success it binds the returning data with result object and return the object
	// If error, return error
	err = stmt.QueryRow(
		transaction.OperationID,
		transaction.AccountID,
		transaction.Amount,
	).Scan(
		&result.ID,
		&result.OperationID,
		&result.AccountID,
		&result.Amount,
		&result.CreatedAt,
	)
	if err != nil {
		slog.Error("error execute the transaction insert statement", slog.String("error", err.Error()))
		return result, err
	}

	return result, nil
}

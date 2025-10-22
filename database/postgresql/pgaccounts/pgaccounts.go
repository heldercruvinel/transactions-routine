package pgaccounts

import (
	"database/sql"
	"log/slog"

	"github.com/heldercruvinel/transactions-routine/internal/accounts"
)

type DB struct {
	db *sql.DB
}

var _ accounts.DB = &DB{}

func GetDB(sqlDB *sql.DB) *DB {
	return &DB{
		db: sqlDB,
	}
}

// PostgreDB method to insert new accounts
func (pg *DB) Insert(account accounts.Account) (accounts.Account, error) {
	var result accounts.Account

	// Create new statement
	stmt, err := pg.db.Prepare("INSERT INTO transactions.accounts (account_code, created_at) VALUES ($1, now()) RETURNING id, account_code, created_at")
	if err != nil {
		slog.Error("error to create account insert statement", slog.String("error", err.Error()))
		return result, err
	}
	defer stmt.Close()

	// Execute the statement, if success it binds the returning data with result object and return the object
	// If error, return error
	err = stmt.QueryRow(
		account.AccountCode,
	).Scan(
		&result.Id,
		&result.AccountCode,
		&result.CreatedAt,
	)
	if err != nil {
		slog.Error("error execute the account insert statement", slog.String("error", err.Error()))
		return result, err
	}

	return result, nil
}

// PostgreDB method to check if the account exists
func (pg *DB) Exists(account accounts.Account) (accounts.Account, error) {
	var result accounts.Account

	// Create new statement
	stmt, err := pg.db.Prepare(`
		SELECT 
			id, 
			account_code, 
			created_at
		FROM transactions.accounts
		WHERE account_code = $1
	`)
	if err != nil {
		slog.Error("error to create account exists statement", slog.String("error", err.Error()))
		return result, err
	}
	defer stmt.Close()

	// Execute the statement, if success it binds the returning data with result object and return the object
	// If error, return error
	err = stmt.QueryRow(
		account.AccountCode,
	).Scan(
		&result.Id,
		&result.AccountCode,
		&result.CreatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		slog.Error("error to execute the account exists statement", slog.String("error", err.Error()))
		return result, err
	}

	return result, nil
}

// PostgreDB method to get a account for account id
func (pg *DB) Get(accountID string) (accounts.Account, error) {
	var result accounts.Account

	// Create new statement
	stmt, err := pg.db.Prepare(`
		SELECT 
			id, 
			account_code, 
			created_at
		FROM transactions.accounts
		WHERE id = $1
	`)
	if err != nil {
		slog.Error("error to create account get statement", slog.String("error", err.Error()))
		return result, err
	}
	defer stmt.Close()

	// Execute the statement, if success it binds the returning data with result object and return the object
	// If error, return error
	err = stmt.QueryRow(
		accountID,
	).Scan(
		&result.Id,
		&result.AccountCode,
		&result.CreatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		slog.Error("error execute the account get statement", slog.String("error", err.Error()))
		return result, err
	}

	return result, nil
}

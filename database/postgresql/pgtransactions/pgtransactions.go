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
		INSERT INTO transactions.transactions (operation_id, account_id, amount, balance, closed, created_at) VALUES
		($1, $2, $3, $4, $5, now())
		RETURNING id, operation_id, account_id, amount, balance, closed, created_at;
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
		transaction.Balance,
		transaction.Closed,
	).Scan(
		&result.ID,
		&result.OperationID,
		&result.AccountID,
		&result.Amount,
		&result.Balance,
		&result.Closed,
		&result.CreatedAt,
	)
	if err != nil {
		slog.Error("error execute the transaction insert statement", slog.String("error", err.Error()))
		return result, err
	}

	return result, nil
}

// PostgreDB method to list opened transactions
func (pg *DB) List(closed bool) ([]transactions.Transaction, error) {
	var total int

	// Create new statement
	countStmt, err := pg.db.Prepare(`
		SELECT count(1) AS total
		FROM transactions.transactions as T
		where T.closed = $1
			and t.operation_id <> 4;
	`)
	if err != nil {
		slog.Error("error to create transaction insert statement", slog.String("error", err.Error()))
		return nil, err
	}
	defer countStmt.Close()

	// Execute a count to get the total registers to create the exactly sized slice
	countRow, err := countStmt.Query(closed)
	if err != nil {
		slog.Error("error to execute the statement", slog.String("error", err.Error()))
		return nil, err
	}

	for countRow.Next() {
		if countRow.Err() != nil {
			if err != nil {
				slog.Error("error to bind countRow", slog.String("error", err.Error()))
				return nil, err
			}
		}
		countRow.Scan(&total)
	}

	// The was executed before to get the total to allow us to create
	// a correctly sized slice and preserve the Garbage Collector from leaking
	result := make([]transactions.Transaction, total)

	// Create new statement
	stmt, err := pg.db.Prepare(`
		SELECT 
			id, 
			operation_id, 
			account_id, 
			amount, 
			balance, 
			closed,
			created_at
		FROM transactions.transactions as T
		 WHERE T.closed = $1
			AND T.operation_id <> 4
		ORDER BY T.created_at ASC;
	`)
	if err != nil {
		slog.Error("error to create transaction insert statement", slog.String("error", err.Error()))
		return result, err
	}
	defer stmt.Close()

	// Execute the statement, if success it binds the returning data with result array and return the object
	// If error, return error
	rows, err := stmt.Query(closed)
	if err != nil {
		slog.Error("error to execute the statement", slog.String("error", err.Error()))
		return result, err
	}

	for i := range result {

		rows.Next()
		if rows.Err() != nil {
			slog.Error("error to bind rows", slog.String("error", err.Error()))
			return result, err
		}

		// r := transactions.Transaction{}
		err = rows.Scan(
			&result[i].ID,
			&result[i].OperationID,
			&result[i].AccountID,
			&result[i].Amount,
			&result[i].Balance,
			&result[i].Closed,
			&result[i].CreatedAt,
		)
		if err != nil {
			slog.Error("error execute the transaction insert statement", slog.String("error", err.Error()))
			return result, err
		}
	}

	return result, nil
}

// PostgreDB method to insert new transactions
func (pg *DB) Update(transaction transactions.Transaction) error {

	// Create new statement
	stmt, err := pg.db.Prepare(`
		UPDATE transactions.transactions
		SET 
			operation_id=$1, 
			account_id=$2, 
			amount=$3,  
			balance=$4, 
			closed=$5
		WHERE id=$6;
	`)
	if err != nil {
		slog.Error("error to create transaction insert statement", slog.String("error", err.Error()))
		return err
	}
	defer stmt.Close()

	// Execute the statement, if success it binds the returning data with result object and return the object
	// If error, return error
	_, err = stmt.Query(
		transaction.OperationID,
		transaction.AccountID,
		transaction.Amount,
		transaction.Balance,
		transaction.Closed,
		transaction.ID,
	)
	if err != nil {
		slog.Error("error execute the transaction insert statement", slog.String("error", err.Error()))
		return err
	}

	return nil
}

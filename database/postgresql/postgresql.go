package postgresql

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

// Creates and returns the sql connection
func GetConnection() (*sql.DB, error) {
	var result int

	// Set the environment variables and create a connectinon string
	databaseHost := os.Getenv("DATABASE_HOST")
	databaseUser := os.Getenv("DATABASE_USER")
	databasePassword := os.Getenv("DATABASE_PASSWORD")

	// TODO: RESOLVER ESSE PROBLEMA DO HOST POSTGRESQL

	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=financial sslmode=disable", databaseUser, databasePassword, databaseHost)

	// Open a connection with the database and return the *sql.DB instance
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		slog.Error("error to connect with database", slog.String("error", err.Error()))
		return nil, err
	}

	// Checking if the "financial" database was created correctly
	err = db.QueryRow("SELECT 1 FROM pg_database WHERE datname = 'financial'").Scan(&result)
	if err != nil || result != 1 {
		slog.Error("error to check if the database exists", slog.String("error", err.Error()))
		return nil, err
	}

	// Checking if the database structure has been created correctly
	// If not, will be created
	err = checkDatabaseStructure(db)
	if err != nil {
		slog.Error("error to check the database structure", slog.String("error", err.Error()))
		return nil, err
	}

	return db, nil
}

// Checks if the database structure has been created
// If not, create it
func checkDatabaseStructure(db *sql.DB) error {

	var exists int

	// Check if the "transactions" schema allready exists
	err := db.QueryRow("SELECT 1 FROM information_schema.schemata WHERE schema_name = 'transactions'").Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		slog.Error("failed when to check the database structure", slog.String("error", err.Error()))
		return err
	}

	// If the schema not exists, will be created
	if exists != 1 {
		_, err = db.Exec(`

			CREATE SCHEMA transactions;

			CREATE TABLE financial.transactions.operations_types (
				id INT4 GENERATED ALWAYS AS IDENTITY,
				description VARCHAR(50) NOT NULL,
				CONSTRAINT pk_operations_types PRIMARY KEY(id)
			);

			INSERT INTO financial.transactions.operations_types (description) VALUES 
			('Normal Puchase'),
			('Purchase with installments'),
			('Withdrawal'),
			('Credit voucher');

			CREATE TABLE financial.transactions.accounts (
				id UUID DEFAULT gen_random_uuid(),
				account_code VARCHAR(50) NOT NULL,
				created_at TIMESTAMP DEFAULT NOW() NOT NULL,
				CONSTRAINT pk_accounts PRIMARY KEY(id)
			);

			CREATE TABLE financial.transactions.transactions (
				id UUID DEFAULT gen_random_uuid(),
				operation_id INT4, 
				account_id UUID NOT NULL,
				amount DECIMAL(10,2) NOT NULL,
				created_at TIMESTAMP DEFAULT NOW() NOT NULL,
				CONSTRAINT pk_transactions PRIMARY KEY(id)
			);

			ALTER TABLE financial.transactions.transactions
			ADD CONSTRAINT fk_operation_id FOREIGN KEY (operation_id) REFERENCES financial.transactions.operations_types (id),
			ADD CONSTRAINT fk_account_id FOREIGN KEY (account_id) REFERENCES financial.transactions.accounts (id);
		`)
		if err != nil {
			slog.Error("failed when to create the database structure", slog.String("error", err.Error()))
			return err
		}
	}

	return nil
}

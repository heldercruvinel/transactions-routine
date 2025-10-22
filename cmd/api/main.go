package main

import (
	"log/slog"
	"net/http"

	"github.com/heldercruvinel/transactions-routine/database/postgresql"
	"github.com/heldercruvinel/transactions-routine/database/postgresql/pgaccounts"
	"github.com/heldercruvinel/transactions-routine/database/postgresql/pgtransactions"
)

func main() {

	db, err := postgresql.GetConnection()
	if err != nil {
		slog.Error("starting server failed", slog.String("error", err.Error()))
		panic(err)
	}

	pgaccounts := pgaccounts.GetDB(db)
	pgtransactions := pgtransactions.GetDB(db)

	server := http.NewServeMux()

	server.HandleFunc("GET /health/{$}", healthHandler(db))
	server.HandleFunc("POST /accounts/{$}", insertAccountHandler(pgaccounts))
	server.HandleFunc("GET /accounts/{accountId}/{$}", getAccountHandler(pgaccounts))
	server.HandleFunc("POST /transactions/{$}", insertTransactionHandler(pgtransactions))

	if err := http.ListenAndServe(":8080", server); err != nil {
		slog.Error("error to start the server", slog.String("error", err.Error()))
	}
}

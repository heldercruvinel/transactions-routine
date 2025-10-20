package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/heldercruvinel/transactions-routine/database/postgresql"
)

func main() {

	db, err := postgresql.GetConnection()
	if err != nil {
		slog.Error("starting server failed", slog.String("error", err.Error()))
		panic(err)
	}

	server := http.NewServeMux()

	server.HandleFunc("GET /health/{$}", healthHandler(db))

	if err := http.ListenAndServe(":8080", server); err != nil {
		slog.Error("error to start the server", slog.String("error", err.Error()))
	}
}

func healthHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbStatus := "on"

		err := db.Ping()
		if err != nil {
			dbStatus = "off"
		}

		fmt.Fprintf(w, "dbStatus: %s", dbStatus)
	}
}

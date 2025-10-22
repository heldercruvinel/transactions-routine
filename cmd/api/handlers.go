package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/heldercruvinel/transactions-routine/internal/accounts"
	"github.com/heldercruvinel/transactions-routine/internal/transactions"
)

// Handle to check api and database health
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

// Handler to inser new accounts
func insertAccountHandler(db accounts.DB) http.HandlerFunc {

	// Using the composition concept
	return func(w http.ResponseWriter, r *http.Request) {
		var account accounts.Account

		// Decoding the json body to a Account object
		err := json.NewDecoder(r.Body).Decode(&account)
		if err != nil {
			slog.Error("error to convert account json into struct", slog.String("error", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error to conver json %v", err.Error())
			return
		}

		// Trys to insert, if the method find any tuple in the database with same value
		// Returns Bad Request Status Code
		// If any other error, returns Internal Server Error Status Code
		result, err := accounts.Insert(account, db)
		if err != nil {
			slog.Error("error to convert account json into struct", slog.String("error", err.Error()))

			if result.Id != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}

			fmt.Fprintf(w, "error to create new account %v", err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result)
	}
}

// Handler to get existing accounts
func getAccountHandler(db accounts.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get the path param accountId
		accountId := r.PathValue("accountId")

		// Search the account inside the database
		// If any error happen, will return Internal Server Error Status Code
		result, err := accounts.Get(accountId, db)
		if err != nil {
			slog.Error("error to get the accountt", slog.String("error", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error to get the accountt %v", err.Error())
			return
		}

		// If none error happens, and find no accounts, return Ok Status Code
		if result.Id == nil {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

// Handler to insert new Transactions
func insertTransactionHandler(db transactions.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction transactions.Transaction

		// Decoding the json body to a Transaction object
		err := json.NewDecoder(r.Body).Decode(&transaction)
		if err != nil {
			slog.Error("error to convert transaction json into struct", slog.String("error", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error to conver json %v", err.Error())
			return
		}

		// Insert the transaction
		// If any error to occur returns Internal Server Error Status Code
		result, err := transactions.Insert(transaction, db)
		if err != nil {
			slog.Error("error to insert transaction", slog.String("error", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error to insert transaction %v", err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result)
	}
}

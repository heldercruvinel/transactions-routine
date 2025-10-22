package transactions

import (
	"log/slog"
	"time"

	"github.com/go-playground/validator/v10"
)

type DB interface {
	Insert(transaction Transaction) (Transaction, error)
}

type Transaction struct {
	ID          *string    `json:"id"`
	OperationID int        `json:"operation_id" validate:"number,required"`
	AccountID   string     `json:"account_id" validate:"uuid,required"`
	Amount      float64    `json:"amount" validate:"numeric,required"`
	CreatedAt   *time.Time `json:"created_at"`
}

var amountRules map[int]func(amount float64) float64 = map[int]func(amount float64) float64{
	1: func(amount float64) float64 {
		if amount > 0 {
			return amount * -1
		}
		return amount
	},
	2: func(amount float64) float64 {
		if amount > 0 {
			return amount * -1
		}
		return amount
	},
	3: func(amount float64) float64 {
		if amount > 0 {
			return amount * -1
		}
		return amount
	},
	4: func(amount float64) float64 {
		if amount < 0 {
			return amount * -1
		}
		return amount
	},
}

// Method responsible for create new Transactions
func Insert(transaction Transaction, db DB) (Transaction, error) {

	// Transaction fields validation
	validate := validator.New(validator.WithRequiredStructEnabled())
	validateErr := validate.Struct(transaction)
	if validateErr != nil {
		slog.Error("error to validate new transaction fields", slog.String("error", validateErr.Error()))
		return Transaction{}, validateErr
	}

	// Validate and apply the signal rules
	transaction.Amount = amountRules[transaction.OperationID](transaction.Amount)

	// Database Transaction insert
	result, err := db.Insert(transaction)
	if err != nil {
		slog.Error("error to create new transacation", slog.String("error", err.Error()))
		return Transaction{}, err
	}

	return result, nil
}

package transactions

import (
	"errors"
	"log/slog"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/heldercruvinel/transactions-routine/numbers"
)

type DB interface {
	Insert(transaction Transaction) (Transaction, error)
	List(closed bool) ([]Transaction, error)
	Update(transaction Transaction) error
}

type Transaction struct {
	ID          *string    `json:"id"`
	OperationID int        `json:"operation_id" validate:"number,required"`
	AccountID   string     `json:"account_id" validate:"uuid,required"`
	Amount      float64    `json:"amount" validate:"numeric,required"`
	Balance     float64    `json:"balance"`
	Closed      bool       `json:"closed"`
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

	result, err := CalcBalance(transaction, db)
	if err != nil {
		slog.Error("error to calcualte positive balance", slog.String("error", err.Error()))
		return Transaction{}, err
	}

	return result, nil

}

// Method responsible for calc the balance and update the opened transactions
func CalcBalance(transaction Transaction, db DB) (Transaction, error) {

	// Validate and apply the signal rules
	rule, ok := amountRules[transaction.OperationID]
	if !ok {
		slog.Error("error to find operation rule")
		return Transaction{}, errors.New("error to find operation rule")
	}
	transaction.Amount = rule(transaction.Amount)

	if transaction.OperationID < 4 {
		transaction.Balance = transaction.Amount
		transaction.Closed = false
	} else {
		// List all tuples that closed is false and OperationID <> 4
		initialBalance := transaction.Amount

		transctionList, err := db.List(false)
		if err != nil {
			slog.Error("error to calcualte positive balance", slog.String("error", err.Error()))
			return Transaction{}, err
		}

		for _, t := range transctionList {

			// Partial disccount
			if (initialBalance + t.Balance) <= 0 {
				t.Balance = numbers.FormatToTwoDecimalPlaces(initialBalance + t.Balance)
				transaction.Balance = 0
				transaction.Closed = true
				initialBalance = 0
			} else {
				// Total disccount
				initialBalance = numbers.FormatToTwoDecimalPlaces(initialBalance + t.Balance)
				t.Closed = true
				t.Balance = 0
			}

			err := db.Update(t)
			if err != nil {
				slog.Error("error to update old balances", slog.String("error", err.Error()))
				return Transaction{}, err
			}
		}

		transaction.Balance = numbers.FormatToTwoDecimalPlaces(initialBalance)
	}

	result, err := db.Insert(transaction)
	if err != nil {
		slog.Error("error to insert positive balance", slog.String("error", err.Error()))
		return Transaction{}, err
	}

	return result, nil
}

package accounts

import (
	"errors"
	"log/slog"
	"time"

	"github.com/go-playground/validator/v10"
)

type DB interface {
	Insert(account Account) (Account, error)
	Exists(account Account) (Account, error)
	Get(accountId string) (Account, error)
}

type Account struct {
	Id          *string    `json:"id"`
	AccountCode string     `json:"account_code" validate:"required,alphanum,max=11,min=11"`
	CreatedAt   *time.Time `json:"created_at"`
}

// Method responsible for create new Accounts
func Insert(account Account, db DB) (Account, error) {

	// Account fields validation
	validate := validator.New(validator.WithPrivateFieldValidation())
	validateErr := validate.Struct(account)
	if validateErr != nil {
		slog.Error("error to validate fields", slog.String("error", validateErr.Error()))
		return Account{}, validateErr.(validator.ValidationErrors)
	}

	// First check if exists any account with same code
	found, err := db.Exists(account)
	if err != nil {
		slog.Error("error to insert account", slog.String("error", err.Error()))
		return found, err
	}

	// In case exists any account with same code, return this Account with a message of error
	// informing that already exists a Account with this code
	if found.Id != nil {
		slog.Error("account already exists")
		return found, errors.New("account already exists")
	}

	// If not exists, do the insert
	result, err := db.Insert(account)
	if err != nil {
		slog.Error("error to insert account", slog.String("error", err.Error()))
		return result, err
	}

	return result, nil
}

// Method responsible for return the found Account
func Get(accountId string, db DB) (Account, error) {

	// Validate if the accountId is a correct uuid and not empty
	validate := validator.New(validator.WithPrivateFieldValidation())
	validateErr := validate.Var(accountId, "uuid,required")
	if validateErr != nil {
		slog.Error("error to validate fields", slog.String("error", validateErr.Error()))
		return Account{}, validateErr.(validator.ValidationErrors)
	}

	// Search at the database for a Account with this ID
	result, err := db.Get(accountId)
	if err != nil {
		slog.Error("error to find a account", slog.String("error", err.Error()))
		return Account{}, err
	}

	return result, nil
}

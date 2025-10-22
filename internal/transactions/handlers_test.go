package transactions

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestInsertTransactionSuccess(t *testing.T) {

	mockedDB := GetMockedDB("")

	t.Run("Insert Transaction with negative value sucess", func(t *testing.T) {
		transaction := Transaction{
			OperationID: 1,
			AccountID:   "dbfc23f8-ffd2-4aaf-ba01-e6eb1213bfbd",
			Amount:      -50.00,
		}

		result, err := Insert(transaction, &mockedDB)
		if err != nil {
			t.Fatal("Should not return error")
		}

		assert.NotEmpty(t, result)
		assert.Equal(t, transaction.OperationID, result.OperationID)
		assert.Equal(t, transaction.AccountID, result.AccountID)
		assert.Equal(t, transaction.Amount, result.Amount)
	})

	t.Run("Insert Transaction with positive value sucess", func(t *testing.T) {
		transaction := Transaction{
			OperationID: 4,
			AccountID:   "dbfc23f8-ffd2-4aaf-ba01-e6eb1213bfbd",
			Amount:      150.00,
		}

		result, err := Insert(transaction, &mockedDB)
		if err != nil {
			t.Fatal("Should not return error")
		}

		assert.NotEmpty(t, result)
		assert.Equal(t, transaction.OperationID, result.OperationID)
		assert.Equal(t, transaction.AccountID, result.AccountID)
		assert.Equal(t, transaction.Amount, result.Amount)
	})

	t.Run("Insert Transaction with positive should transform into negative value sucess", func(t *testing.T) {
		transaction := Transaction{
			OperationID: 1,
			AccountID:   "dbfc23f8-ffd2-4aaf-ba01-e6eb1213bfbd",
			Amount:      150.00,
		}

		result, err := Insert(transaction, &mockedDB)
		if err != nil {
			t.Fatal("Should not return error")
		}

		assert.NotEmpty(t, result)
		assert.Equal(t, transaction.OperationID, result.OperationID)
		assert.Equal(t, transaction.AccountID, result.AccountID)
		assert.Equal(t, transaction.Amount*-1, result.Amount)
	})
}

func TestInsertTransactionError(t *testing.T) {

	t.Run("Should return database error", func(t *testing.T) {
		inserError := "error database connection"
		mockedDB := GetMockedDB(inserError)

		transaction := Transaction{
			OperationID: 1,
			AccountID:   "8403fd34-eb6f-45cf-8724-d0f5abdee42d",
			Amount:      75.00,
		}

		result, err := Insert(transaction, &mockedDB)
		if err == nil {
			t.Fatal("Should return error")
		}

		assert.Empty(t, result)
		assert.Equal(t, inserError, err.Error())
	})

	t.Run("Should return AccountID not equal uuid validate error", func(t *testing.T) {
		mockedDB := GetMockedDB("")

		transaction := Transaction{
			OperationID: 1,
			AccountID:   "asdfasdfasdfasdfasdf",
			Amount:      150.00,
		}

		result, err := Insert(transaction, &mockedDB)
		if err == nil {
			t.Fatal("Should return error")
		}

		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			t.Errorf("Expected ValidationErrors, got %T", err)
		}

		assert.Empty(t, result)
		assert.Equal(t, "AccountID", validationErrors[0].Field())
	})

	t.Run("Should return validate OperationID nil error", func(t *testing.T) {
		mockedDB := GetMockedDB("")

		transaction := Transaction{
			AccountID: "dbfc23f8-ffd2-4aaf-ba01-e6eb1213bfbd",
			Amount:    150.00,
		}

		result, err := Insert(transaction, &mockedDB)
		if err == nil {
			t.Fatal("Should return error")
		}

		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			t.Errorf("Expected ValidationErrors, got %T", err)
		}

		assert.Empty(t, result)
		assert.Equal(t, "OperationID", validationErrors[0].Field())
	})

	t.Run("Should return validate Amount nil error", func(t *testing.T) {
		mockedDB := GetMockedDB("")

		transaction := Transaction{
			OperationID: 1,
			AccountID:   "dbfc23f8-ffd2-4aaf-ba01-e6eb1213bfbd",
		}

		result, err := Insert(transaction, &mockedDB)
		if err == nil {
			t.Fatal("Should return error")
		}

		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			t.Errorf("Expected ValidationErrors, got %T", err)
		}

		assert.Empty(t, result)
		assert.Equal(t, "Amount", validationErrors[0].Field())
	})
}

package accounts

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestInsertAccountsSuccess(t *testing.T) {

	mockedDB := GetMockedDB("", "", "")

	t.Run("Insert account 0000998772x should return success", func(t *testing.T) {

		newAccount := Account{
			AccountCode: "0000998772x",
		}

		result, err := Insert(newAccount, &mockedDB)
		if err != nil {
			t.Error("Expected success but got error")
		}

		if !assert.Equal(t, newAccount.AccountCode, result.AccountCode) && result.Id != nil {
			t.Errorf("Expected: %v but got %v", newAccount, result)
		}
	})

	t.Run("Insert account 00109877376x should return success", func(t *testing.T) {

		newAccount := Account{
			AccountCode: "0010987776x",
		}

		result, err := Insert(newAccount, &mockedDB)
		if err != nil {
			t.Error("Expected success but got error")
		}

		if !assert.Equal(t, newAccount.AccountCode, result.AccountCode) && result.Id != nil {
			t.Errorf("Expected: %v but got %v", newAccount, result)
		}
	})
}

func TestInsertAccountsErrors(t *testing.T) {

	t.Run("Insert account 0000999992x should return error at Insert", func(t *testing.T) {

		mockedDB := GetMockedDB("database connection error", "", "")

		newAccount := Account{
			AccountCode: "0000999992x",
		}

		result, err := Insert(newAccount, &mockedDB)
		if err == nil {
			t.Error("Expected success but got error")
		}

		if !assert.Equal(t, "database connection error", err.Error()) && result.Id != nil {
			t.Errorf("Expected: %v but got %v", newAccount, result)
		}
	})

	t.Run("Insert account 0000333992x should return error at Exists", func(t *testing.T) {

		mockedDB := GetMockedDB("", "database connection error", "")

		newAccount := Account{
			AccountCode: "0000333992x",
		}

		result, err := Insert(newAccount, &mockedDB)
		if err == nil {
			t.Error("Expected success but got error")
		}

		if !assert.Equal(t, "database connection error", err.Error()) && result.Id != nil {
			t.Errorf("Expected: %v but got %v", newAccount, result)
		}
	})

	t.Run("Insert account 00003339922x should return error at Validator", func(t *testing.T) {

		mockedDB := GetMockedDB("", "", "")

		newAccount := Account{
			AccountCode: "00003339922x",
		}

		result, err := Insert(newAccount, &mockedDB)
		if err == nil {
			t.Error("Expected success but got error")
		}

		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			t.Errorf("Expected ValidationErrors, got %T", err)
		}

		if !assert.Equal(t, "AccountCode", validationErrors[0].Field()) && result.Id != nil {
			t.Errorf("Expected: %v but got %v", "AccountCode", validationErrors[0].Field())
		}
	})

	t.Run("Insert account 0000333392x should return error when to find the Account", func(t *testing.T) {

		mockedDB := GetMockedDB("", "found", "")

		newAccount := Account{
			AccountCode: "0000333392x",
		}

		result, err := Insert(newAccount, &mockedDB)
		if err == nil {
			t.Error("Expected success but got error")
		}

		if !assert.Equal(t, "account already exists", err.Error()) && result.Id != nil {
			t.Errorf("Expected: %v but got %v", "AccountCode", err.Error())
		}
	})
}

func TestGetAccountSuccess(t *testing.T) {

	mockedDB := GetMockedDB("", "", "")

	t.Run("Should return a account with success", func(t *testing.T) {

		accountId := "83064af3-bb81-4514-a6d4-afba340825cd"
		result, err := Get(accountId, &mockedDB)
		if err != nil {
			t.Fatal("Should not return error")
		}

		assert.NotEmpty(t, result)
		assert.Equal(t, &accountId, result.Id)
	})

	t.Run("Should not return a Account with success", func(t *testing.T) {

		accountId := "6cc1f11d-5f58-4019-b53e-de84f66c579b"
		result, err := Get(accountId, &mockedDB)
		if err != nil {
			t.Fatal("Should not return error")
		}

		assert.Empty(t, result)
	})
}

func TestGetAccountError(t *testing.T) {

	t.Run("Should not return a account with database error", func(t *testing.T) {

		errText := "database connection"

		mockedDB := GetMockedDB("", "", errText)

		accountId := "83064af3-bb81-4514-a6d4-afba340825cd"
		result, err := Get(accountId, &mockedDB)
		if err == nil {
			t.Fatal("Should return error")
		}

		assert.Empty(t, result)
		assert.Equal(t, errText, err.Error())
	})

	t.Run("Should not return a Account with validate error", func(t *testing.T) {

		mockedDB := GetMockedDB("", "", "")

		accountId := "adfasdfasdfasdasdadsf"
		result, err := Get(accountId, &mockedDB)
		if err == nil {
			t.Fatal("Should not return error")
		}

		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			t.Errorf("Expected ValidationErrors, got %T", err)
		}

		assert.Empty(t, result)
		assert.Equal(t, "uuid", validationErrors[0].Tag())
	})
}

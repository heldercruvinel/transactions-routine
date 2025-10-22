package accounts

import (
	"errors"
	"time"
)

var accountRepo []Account = []Account{
	{Id: func() *string { v := "83064af3-bb81-4514-a6d4-afba340825cd"; return &v }(), AccountCode: "0000988772x"},
}

type MockDB struct {
	funcInsert func(a Account) (Account, error)
	funcExists func(a Account) (Account, error)
	funcGet    func(id string) (Account, error)
}

var _ DB = &MockDB{}

func GetMockedDB(
	errorInsert string,
	errorExists string,
	errorGet string,
) MockDB {
	mockedDB := MockDB{
		funcInsert: funcInsert(errorInsert),
		funcExists: funcExists(errorExists),
		funcGet:    funcGet(errorGet),
	}

	return mockedDB
}

func (db *MockDB) Insert(account Account) (Account, error) {
	return db.funcInsert(account)
}

func (db *MockDB) Exists(account Account) (Account, error) {
	return db.funcExists(account)
}

func (db *MockDB) Get(id string) (Account, error) {
	return db.funcGet(id)
}

func funcInsert(errorInsert string) func(a Account) (Account, error) {
	if errorInsert != "" {
		return func(a Account) (Account, error) {
			return Account{}, errors.New(errorInsert)
		}
	}

	return func(a Account) (Account, error) {
		return Account{
			Id:          func() *string { v := "uuuuuuuuuiiiiiiiiiddddddd"; return &v }(),
			AccountCode: a.AccountCode,
			CreatedAt:   func() *time.Time { v := time.Now(); return &v }(),
		}, nil
	}
}

func funcExists(errorExists string) func(a Account) (Account, error) {
	if errorExists != "" && errorExists != "found" {
		return func(a Account) (Account, error) {
			return Account{}, errors.New(errorExists)
		}
	}

	if errorExists == "found" {
		return func(a Account) (Account, error) {
			return Account{
				Id:          func() *string { v := "lkajdfoiasudfjasdiojuasdfjsa"; return &v }(),
				AccountCode: "00098623410x",
			}, nil
		}
	}

	return func(a Account) (Account, error) {
		for _, item := range accountRepo {
			if item.AccountCode == a.AccountCode {
				return item, nil
			}
		}

		return Account{}, nil
	}
}

func funcGet(errorGet string) func(id string) (Account, error) {
	if errorGet != "" {
		return func(id string) (Account, error) {
			return Account{}, errors.New(errorGet)
		}
	}

	return func(id string) (Account, error) {
		for _, item := range accountRepo {
			if *item.Id == id {
				return item, nil
			}
		}

		return Account{}, nil
	}
}

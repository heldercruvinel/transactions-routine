package transactions

import (
	"errors"
	"time"
)

var transactionRepo []Transaction = []Transaction{
	{
		ID:          func() *string { v := "f1f7f26a-897e-4cc0-b30f-3a4dba00fcc4"; return &v }(),
		OperationID: 1,
		AccountID:   "4bf92821-1117-4fcf-aeef-55070ecca3bd",
		Amount:      150.00,
		CreatedAt:   func() *time.Time { v := time.Now(); return &v }(),
	},
}

type MockDB struct {
	funcInsert func(a Transaction) (Transaction, error)
}

var _ DB = &MockDB{}

func GetMockedDB(
	errorInsert string,
) MockDB {
	mockedDB := MockDB{
		funcInsert: funcInsert(errorInsert),
	}

	return mockedDB
}

func (db *MockDB) Insert(account Transaction) (Transaction, error) {
	return db.funcInsert(account)
}

func funcInsert(errorInsert string) func(a Transaction) (Transaction, error) {
	if errorInsert != "" {
		return func(a Transaction) (Transaction, error) {
			return Transaction{}, errors.New(errorInsert)
		}
	}

	return func(a Transaction) (Transaction, error) {
		transactionRepo = append(transactionRepo, a)

		return Transaction{
			OperationID: a.OperationID,
			AccountID:   a.AccountID,
			Amount:      a.Amount,
		}, nil
	}
}

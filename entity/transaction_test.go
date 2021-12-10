package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Criando função GoLang
//Testando transação maior que 1000
func TestTransactionWithAmountGreaterThan1000(t *testing.T) {
	// ":=" Criando e declarando váriavel
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 2000
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "you dont have limit for this transaction", err.Error())
}

func TestTransactionWithAmountlesserThan1(t *testing.T) {
	// ":=" Criando e declarando váriavel
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "the amount must br greater than one(1)", err.Error())
}

package entity

import (
	"errors"
)

//Classe (c/ propriedades e métodos)

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("you dont have limit for this transaction")
	}
	if t.Amount < 1 {
		return errors.New("the amount must br greater than one(1)")
	}
	return nil
}

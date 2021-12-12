package process_transaction

import "github.com/Alladio/GoLang/entity"

type ProcessTrasction struct {
	Repository entity.TransactionRepository
}

func NewProcessTransaction(repository entity.TransactionRepository) *ProcessTrasction {
	return &ProcessTrasction{Repository: repository}
}

func (p *ProcessTrasction) Execute(input TransactionDtoInput) error {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	invalidTransaction := transaction.IsValid()
	if invalidTransaction != nil {
		err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, transaction.Status)
		if err != nil {
			return Trans
		}
	}

	return nil
}

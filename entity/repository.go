package entity

type TransactionRepository interface {
	Insert(id string, accountId string, ammount float64, status string, errorMessage string) error
}

package entity

type TransactionRepository interface {
	Insert(id string, accouintId string, amount float64, status string, errorMessage string) error
}
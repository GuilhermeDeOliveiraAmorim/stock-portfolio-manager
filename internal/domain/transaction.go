package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID string `json:"id"`
	StockID string `json:"stock_id"`
	TransactionType string `json:"type"`
	AddedIn string `json:"added_in"`
	DeactivatedIn string `json:"deactivated_in"`
}

func NewTransaction(stockID string, transactiontype string) (*Transaction, error) {
	dateNow := time.Now().Local().String()
	transaction := &Transaction{
		ID: uuid.New().String(),
		StockID: stockID,
		TransactionType: transactiontype,
		AddedIn: dateNow,
		DeactivatedIn: "",
	}

	return transaction, nil
}
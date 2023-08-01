package domain

import (
	"time"

	"github.com/google/uuid"
)

type Stock struct {
	ID string `json:"id"`
	Ticker string `json:"ticker"`
	ImageID string `json:"image_id"`
	MarketCap float64 `json:"market_cap"`
	PL float64 `json:"pl"`
	Yield float64 `json:"yield"`
	Transactions []*Transaction `json:"transactions"`
	CreatedAt string
}

func NewStock(ticker string, imageId string, marketCap float64, pl float64, yield float64) (*Stock, error) {
	dateNow := time.Now().Local().String()
	stock := &Stock{
		ID: uuid.New().String(),
		Ticker: ticker,
		ImageID: imageId,
		MarketCap: marketCap,
		PL: pl,
		Yield: yield,
		CreatedAt: dateNow,
	}

	return stock, nil
}

func (stock *Stock) AddTransaction(transactionToAdd *Transaction) {
	stock.Transactions = append(stock.Transactions, transactionToAdd)
}
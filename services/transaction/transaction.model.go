package transaction

import (
	"gorm.io/gorm"
	"portfoyum/utils/database"
	"time"
)

type Transaction struct {
	gorm.Model
	Stock      string    `json:"stock"`
	Date       time.Time `json:"date"`
	Quantity   int       `json:"quantity"`
	Price      float64   `json:"price"`
	StockPrice float64   `json:"stock_price"`
	Commission float64   `json:"commission"`
	Type       string    `json:"type"`
	Duplicated bool      `json:"duplicated" gorm:"-"`
	Import     bool      `json:"import" gorm:"-"`
}

func CreateTransaction(t *Transaction) *gorm.DB {
	return database.DB.Create(t)
}
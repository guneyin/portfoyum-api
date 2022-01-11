package transaction

import (
	"gorm.io/gorm"
	"portfoyum-api/services/stock"
	"portfoyum-api/utils/database"
	"time"
)

type Transaction struct {
	SymbolCode string       `gorm:"size:20" json:"code"`
	Date       time.Time    `json:"date"`
	Quantity   int          `json:"quantity"`
	Price      float64      `json:"price"`
	StockPrice float64      `json:"stock_price"`
	Commission float64      `json:"commission"`
	Type       string       `gorm:"size:30" json:"type"`
	Duplicated bool         `json:"duplicated" gorm:"-"`
	Import     bool         `json:"import" gorm:"-"`
	Symbol     stock.Symbol `gorm:"foreignKey:SymbolCode;references:Code" json:"symbol"`
}

func CreateTransaction(t *Transaction) *gorm.DB {
	return database.DB.Create(t)
}

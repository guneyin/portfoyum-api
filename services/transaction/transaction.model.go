package transaction

import (
	"gorm.io/gorm"
	"portfoyum-api/services/stock"
	"portfoyum-api/utils/database"
	"time"
)

type Transaction struct {
	UserID     uint         `json:"user_id"`
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
	//Compared   Compared     `json:"compared" gorm:"-"`
}

type Compared struct {
	Symbol     string  `json:"symbol"`
	Price      float64 `json:"price"`
	TotalPrice float64 `json:"total_price"`
}

func CreateTransaction(t *Transaction) *gorm.DB {
	return database.DB.Create(t)
}

func (c *Compared) Init(d time.Time, q int) {
	if c.Symbol == "" {
		return
	}

	er := new(stock.ExchangeRate)

	c.Price = er.Get(c.Symbol, d)
	c.TotalPrice = c.Price * float64(q)
}

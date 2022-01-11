package stock

import (
	"github.com/piquette/finance-go"
	"portfoyum-api/utils/api"
	"portfoyum-api/utils/database"
	"time"
)

type Symbol struct {
	Code string `gorm:"primaryKey;not null;size:20" json:"code"`
	Name string `gorm:"not null;size:250" json:"name"`
	Slug string `gorm:"size:250" json:"slug"`
}

type SymbolList struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		S string   `json:"s"`
		D []string `json:"d"`
	} `json:"data"`
}

type Equity struct {
	Code string `gorm:"primaryKey;not null;size:20" json:"code"`
	finance.Equity
}

type ExchangeRate struct {
	Symbol string    `gorm:"primaryKey;not null;size:5" json:"symbol"`
	Date   time.Time `gorm:"primaryKey;not null;type:date" json:"date"`
	Rate   float64   `gorm:"not null" json:"rate"`
}

func (er *ExchangeRate) Get(symbol string, date time.Time) float64 {
	er.Symbol = symbol
	er.Date = date

	r := database.DB.First(&er)

	if r.RowsAffected == 0 {
		p := api.GetExchangeRate(er.Symbol, er.Date)

		if p > 0 {
			er.Rate = p

			database.DB.Debug().Save(&er)
		}
	}

	return er.Rate
}

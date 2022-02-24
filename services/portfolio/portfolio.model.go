package portfolio

import (
	"fmt"
	"portfoyum-api/services/stock"
	"portfoyum-api/services/transaction"
	"portfoyum-api/utils/database"
)

type Item struct {
	SymbolCode string               `json:"symbol_code"`
	Quantity   int                  `json:"quantity"`
	TotalBuy   float64              `json:"price"`
	Symbol     stock.Symbol         `json:"symbol"`
	Compared   transaction.Compared `json:"compared"`
}

type Portfolio struct {
	Items []struct {
		SymbolCode string       `json:"symbol_code"`
		Price      float64      `json:"price"`
		Quantity   int          `json:"quantity"`
		TotalBuy   float64      `json:"total_buy"`
		TotalPrice float64      `json:"total_price"`
		Symbol     stock.Symbol `json:"symbol"`
		//Compared   transaction.Compared `json:"compared"`
	} `json:"items"`
}

func (p *Portfolio) Init(Symbol string) {
	symbolCondition := ""

	if Symbol != "" {
		symbolCondition = fmt.Sprintf("where s.symbol_code = '%s'", Symbol)
	}

	database.DB.Raw(
		"select t.*, t.price * t.quantity as total_price " +
			"from  (" +
			"select s.symbol_code, e.regular_market_price as price, sum(s.quantity) quantity, sum(s.price) total_buy " +
			"from ( " +
			"select symbol_code, " +
			"case " +
			"when type = 'Hisse Satış' then -1 * quantity " +
			"else quantity " +
			"end quantity, " +
			"case " +
			"when type = 'Hisse Satış' then -1 * price " +
			"else price " +
			"end price " +
			"from transactions) as s " +
			"inner join equities e on (e.code = s.symbol_code)" +
			symbolCondition +
			"group by s.symbol_code, e.regular_market_price " +
			"having sum(s.quantity) > 0) as t ").Scan(&p.Items)

	for i, item := range p.Items {
		database.DB.First(&p.Items[i].Symbol, "code = ?", item.SymbolCode)
		//
		//	p.Items[i].Compared.Symbol = symbol
		//
		//	database.DB.Debug().Where("symbol_code = ?", item.SymbolCode).Find(&transactions)
		//
		//	for _, t := range transactions {
		//		t.Compared.Symbol = symbol
		//		t.Compared.Init(t.Date, t.Quantity)
		//
		//		p.Items[i].Compared.Price += t.Compared.Price
		//		p.Items[i].Compared.TotalPrice += t.Compared.TotalPrice
		//	}
		//
		//	p.Items[i].Compared.Price = p.Items[i].Compared.Price / float64(len(transactions))
	}
}

package stats

import (
	"portfoyum-api/utils/database"
)

type Stats struct {
	TotalQty      int     `json:"totalQty"`
	TotalSymbol   int     `json:"totalSymbol"`
	TotalBuy      float64 `json:"totalBuy"`
	TotalPrice    float64 `json:"totalPrice"`
	GainByPrice   float64 `json:"gainByPrice"`
	GainByPercent float64 `json:"gainByPercent""`
}

func (s *Stats) Init() {
	row := database.DB.Table("transactions").Where("type = ?", "Hisse Alış").Select("sum(quantity)").Row()
	row.Scan(&s.TotalQty)

	row = database.DB.Table("transactions").Where("type = ?", "Hisse Alış").Select("count(distinct symbol_code)").Row()
	row.Scan(&s.TotalSymbol)

	row = database.DB.Table("transactions").Where("type = ?", "Hisse Alış").Select("sum(price)").Row()
	row.Scan(&s.TotalBuy)

	row = database.DB.Table("transactions").Select("sum(transactions.quantity * e.regular_market_price) as market_price ").Joins("inner join equities e on e.code = transactions.symbol_code").Row()
	row.Scan(&s.TotalPrice)

	s.GainByPrice = s.TotalPrice - s.TotalBuy

	if s.GainByPrice > 0 {
		s.GainByPercent = (s.GainByPrice / s.TotalBuy) * 100
	} else {
		s.GainByPercent = 0
	}

}

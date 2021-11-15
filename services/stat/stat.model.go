package stat

import "portfoyum-api/utils/database"

type Stats struct {
	TotalQty      int   	`json:"totalQty"`
	TotalBuy	  float64   `json:"totalBuy"`
	TotalPrice    float64   `json:"totalPrice"`
}

func (s *Stats) Init() {
	row := database.DB.Table("transactions").Where("type = ?", "Hisse Alış").Select("sum(quantity)").Row()
	row.Scan(&s.TotalQty)

	row = database.DB.Table("transactions").Where("type = ?", "Hisse Alış").Select("sum(price)").Row()
	row.Scan(&s.TotalBuy)
}

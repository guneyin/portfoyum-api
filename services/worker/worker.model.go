package worker

type StockList struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		S string        `json:"s"`
		D []string `json:"d"`
	} `json:"data"`
}


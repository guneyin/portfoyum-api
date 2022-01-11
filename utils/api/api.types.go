package api

type GetExchangeRateDTO struct {
	Success bool                   `json:"success"`
	Rates   map[string]interface{} `json:"rates"`
}

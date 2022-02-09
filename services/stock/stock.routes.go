package stock

import (
	"github.com/gofiber/fiber/v2"
)

func StockRoutes(app fiber.Router) {
	r := app.Group("/stock")

	r.Get("/equities", GetEquities)
	r.Get("/equities/:code?", GetEquity)

	r.Get("/equities/sync", SyncEquities)
	r.Get("/symbols/sync", SyncSymbols)
	r.Get("exchange/get", GetExchangeRates)
}

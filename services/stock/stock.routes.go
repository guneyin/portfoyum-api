package stock

import (
	"github.com/gofiber/fiber/v2"
)

func StockRoutes(app fiber.Router) {
	r := app.Group("/stock")

	r.Get("/symbols/sync", SyncSymbols)
	r.Get("/symbols/get/:code?", GetSymbols)
	r.Get("/symbols/detail/sync", SyncSymbolsDetail)
	r.Get("/symbols/getList", GetSymbolList)
}

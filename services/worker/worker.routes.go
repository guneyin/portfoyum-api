package worker

import "github.com/gofiber/fiber/v2"

func WorkerRoutes(app fiber.Router) {
	r := app.Group("/worker")

	r.Get("/syncStocks", SyncStocks)
}

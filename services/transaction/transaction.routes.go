package transaction

import "github.com/gofiber/fiber/v2"

func TransactionRoutes(app fiber.Router) {
	r := app.Group("/transactions")

	r.Post("/import", ImportTransactions)
	r.Post("/save", SaveTransactions)
}
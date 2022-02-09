package transaction

import (
	"github.com/gofiber/fiber/v2"
	"portfoyum-api/utils/middleware"
)

func TransactionRoutes(app fiber.Router) {
	r := app.Group("/transactions").Use(middleware.Auth)

	r.Post("/upload", UploadTransactions)
	r.Post("/save", SaveTransactions)
	r.Get("/:symbol?", GetTransactions)
}

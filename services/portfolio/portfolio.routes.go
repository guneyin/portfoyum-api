package portfolio

import (
	"github.com/gofiber/fiber/v2"
	"portfoyum-api/utils/middleware"
)

func PortfolioRoutes(app fiber.Router) {
	r := app.Group("/portfolio").Use(middleware.Auth)

	r.Get("/:symbol?", GetPortfolio)
}

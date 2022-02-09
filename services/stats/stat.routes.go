package stats

import (
	"github.com/gofiber/fiber/v2"
	"portfoyum-api/utils/middleware"
)

func StatRoutes(app fiber.Router) {
	r := app.Group("/stats").Use(middleware.Auth)

	r.Get("/", GetStats)
	//r.Get("/portfolio/:compare?/*", GetPortfolio)
}

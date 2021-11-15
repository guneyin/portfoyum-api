package stat

import "github.com/gofiber/fiber/v2"

func StatRoutes(app fiber.Router) {
	r := app.Group("/stats")

	r.Get("/", GetStats)
}
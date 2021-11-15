package admin

import "github.com/gofiber/fiber/v2"

func AdminRoutes(app fiber.Router) {
	r := app.Group("/admin")

	r.Post("/login", AdminLogin)
}

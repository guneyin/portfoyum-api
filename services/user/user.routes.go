package user

import (
	"github.com/gofiber/fiber/v2"
	"portfoyum-api/utils/middleware"
)

// UserRoutes contains all routes relative to /user
func UserRoutes(app fiber.Router) {
	r := app.Group("/user").Use(middleware.Auth)

	r.Get("/me", UserMe)
	r.Put("/update", UserUpdate)
	r.Delete("/:id", UserDelete)
}

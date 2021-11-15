package auth

import (
	"github.com/gofiber/fiber/v2"
)

// AuthRoutes containes all the auth routes
func AuthRoutes(app fiber.Router) {
	r := app.Group("/auth")

	r.Post("/signup", AuthSignup)
	r.Post("/login", AuthLogin)
	r.Post("/password/forgot", AuthForgotPassword)
	r.Get("/password/verify/:token", AuthVerifyToken)
	r.Post("/password/change", AuthChangePassword)
}

package auth

import (
	"github.com/gofiber/fiber/v2"
)

// AuthRoutes containes all the auth routes
func AuthRoutes(app fiber.Router) {
	r := app.Group("/auth")

	r.Post("/signup", Signup)
	r.Post("/login", Login)
	r.Post("/password/forgot", ForgotPassword)
	r.Get("/verify/:token", VerifyToken)
	r.Post("/password/change", ChangePassword)
}

package main

import (
	"fmt"
	"portfoyum/config"
	"portfoyum/services/admin"
	"portfoyum/services/auth"
	"portfoyum/services/stat"
	"portfoyum/services/stock"
	"portfoyum/services/transaction"
	"portfoyum/services/user"
	"portfoyum/utils"
	"portfoyum/utils/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.Init()

	database.Connect()
	err := database.Migrate(&admin.Admin{}, &user.User{}, &stock.Symbol{}, &stock.SymbolDetail{}, &stock.Equity{}, &transaction.Transaction{})
	if err != nil {
		return
	}

	admin.InitAdmin()

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")

	admin.AdminRoutes(v1)
	auth.AuthRoutes(v1)
	user.UserRoutes(v1)
	stock.StockRoutes(v1)
	transaction.TransactionRoutes(v1)
	stat.StatRoutes(v1)

	app.Listen(fmt.Sprintf(":%v", config.Settings.Server.Port))
}
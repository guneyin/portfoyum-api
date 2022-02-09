package portfolio

import "github.com/gofiber/fiber/v2"

func GetPortfolio(c *fiber.Ctx) error {
	symbol := c.Params("symbol")

	p := new(Portfolio)
	p.Init(symbol)

	return c.JSON(p)
}

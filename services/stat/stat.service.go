package stat

import "github.com/gofiber/fiber/v2"

func GetStats (c *fiber.Ctx) error {
	s := new(Stats)
	s.Init()

	return c.JSON(s)
}

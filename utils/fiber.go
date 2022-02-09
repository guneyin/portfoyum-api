package utils

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseHTTP struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ParseBody(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

func ParseBodyAndValidate(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ParseBody(ctx, body); err != nil {
		return err
	}

	return Validate(body)
}

func GetUserId(c *fiber.Ctx) *uint {
	id, _ := c.Locals("USER").(uint)

	return &id
}

//func Response(message string, data ...interface{}) *ResponseHTTP {
//	return &ResponseHTTP{
//		Message: message,
//		Data:    data[0],
//	}
//}

func Response(c *fiber.Ctx, message string, data ...interface{}) error {
	content := &ResponseHTTP{
		Message: message,
		Data:    data[0],
	}

	return c.JSON(content)
}

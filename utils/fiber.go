package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
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

func GetUserId(c *fiber.Ctx) *uuid.UUID {
	id, _ := c.Locals("USER").(uuid.UUID)

	return &id
}

func Response(message string, data ...interface{}) *ResponseHTTP {
	//if data == nil {
	//	data = make([]interface{}, 0)
	//}

	return &ResponseHTTP{
		Success: true,
		Message: message,
		Data: data,
	}
}

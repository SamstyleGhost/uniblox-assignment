package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CheckUserValidity(c *fiber.Ctx) error {

	if c.Method() != fiber.MethodPost {
		return c.Next()
	}

	payload := struct {
		UserID uuid.UUID `json:"user_id"`
	}{}

	// Would return error if the request body is not set correctly
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	if payload.UserID == uuid.Nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": "Nil uuid received",
		})
	}

	c.Locals("user_id", payload.UserID)

	return c.Next()
}

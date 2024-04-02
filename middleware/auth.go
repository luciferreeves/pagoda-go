package middleware

import (
	"pagoda/auth"

	"github.com/gofiber/fiber/v2"
)

type Session struct {
	SessionId string `cookie:"session"`
}

func Authenticated(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	// check if the header is empty
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	user, err := auth.GetSession(authHeader)

	// user, err := auth.GetSession(session.SessionId)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// set the user in the context
	c.Locals("user", user)

	return c.Next()
}

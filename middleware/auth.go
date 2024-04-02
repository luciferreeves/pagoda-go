package middleware

import (
	"pagoda/auth"

	"github.com/gofiber/fiber/v2"
)

type Session struct {
	SessionId string `cookie:"session"`
}

func Authenticated(c *fiber.Ctx) error {
	session := new(Session)

	if err := c.CookieParser(session); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errorCode": fiber.StatusUnauthorized,
			"error":     "Unauthorized",
		})
	}

	user, err := auth.GetSession(session.SessionId)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errorCode": fiber.StatusUnauthorized,
			"error":     "Unauthorized",
		})
	}

	// set the user in the context
	c.Locals("user", user)

	return c.Next()
}

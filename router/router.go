package router

import (
	"pagoda/controllers"
	"pagoda/middleware"

	"github.com/gofiber/fiber/v2"
)

func Initialize(router *fiber.App) {
	router.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "pong",
			"version": "1.0.0",
		})
	})

	router.Use(middleware.Json)

	router.Use(middleware.Security)

	users := router.Group("/users")

	users.Post("/create", controllers.CreateUser)
	users.Post("/login", controllers.LoginUser)
	users.Post("/logout", middleware.Authenticated, controllers.LogoutUser)
	users.Get("/current", middleware.Authenticated, controllers.CurrentUser)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errorCode": fiber.StatusNotFound,
			"error":     "API Route Not Found",
		})
	})
}

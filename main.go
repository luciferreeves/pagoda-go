package main

import (
	"log"
	"os"
	"pagoda/database"
	"pagoda/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	// Load the .env file
	godotenv.Load()
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	app.Use(logger.New())

	router.Initialize(app)
	database.ConnectDB()
	database.ConnectRedis()

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

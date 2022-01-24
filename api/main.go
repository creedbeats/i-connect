package main

import (
	"log"

	"github.com/creedbeats/i-connect.git/api/config"
	"github.com/creedbeats/i-connect.git/api/database"
	handlers "github.com/creedbeats/i-connect.git/api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK!")
	})
	v1 := app.Group("/v1", logger.New())
	handlers.UserRoutes(v1)
	log.Fatal(app.Listen(":" + config.Get("API_PORT")))
}

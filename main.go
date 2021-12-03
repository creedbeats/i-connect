package main

import (
	"log"

	"github.com/creedbeats/i-connect.git/api/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes (app *fiber.App) {
	app.Use(logger.New())
	app.Get("/v1/users", users.ListUsers)
	app.Post("/v1/users", users.CreateUser)
	app.Get("/v1/users/:id", users.GetUser)
	app.Put("/v1/users/:id", users.UpdateUser)
	app.Delete("/v1/users/:id", users.DeleteUser)
} 

func main() {
	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}

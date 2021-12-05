package api

import (
	"log"

	controllers "github.com/creedbeats/i-connect.git/api/controllers/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	app.Use(logger.New())
	app.Get("/v1/users", controllers.ListUsers)
	app.Post("/v1/users", controllers.CreateUser)
	app.Get("/v1/users/:id", controllers.GetUser)
	app.Put("/v1/users/:id", controllers.UpdateUser)
	app.Delete("/v1/users/:id", controllers.DeleteUser)
}

func Initialize() {
	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
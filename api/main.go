package main

import (
	"log"
	"os"

	controllers "github.com/creedbeats/i-connect.git/api/controllers/user"
	"github.com/creedbeats/i-connect.git/api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/v1/users", controllers.ListUsers)
	app.Post("/v1/users", controllers.CreateUser)
	app.Get("/v1/users/:id", controllers.GetUser)
	app.Put("/v1/users/:id", controllers.UpdateUser)
	app.Delete("/v1/users/:id", controllers.DeleteUser)
	log.Fatal(app.Listen(os.Getenv("API_PORT")))
}

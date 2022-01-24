package handlers

import (
	"github.com/creedbeats/i-connect.git/api/database"
	"github.com/creedbeats/i-connect.git/api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ListUsers(c *fiber.Ctx) error {
	db := database.DB
	user := models.User{}
	users, err := user.List(db)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) (err error) {
	db := database.DB
	user := models.User{}
	if err = c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Create the User and return error if encountered
	if err = user.Create(db); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) (err error) {
	db := database.DB
	user := models.User{}
	user.ID, err = uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	if err = user.Get(db); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "The requested user could not be found", "data": err})
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) (err error) {
	db := database.DB
	user := models.User{}
	
	if user.ID, err = uuid.Parse(c.Params("id")); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "The requested user could not be found", "data": err})
	}
	if err = c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	if err = user.Get(db); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "The requested user could not be found", "data": err})
	}
	
	err = user.Update(db)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) (err error) {
	db := database.DB
	user := models.User{}
	user.ID, err = uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	user.Delete(db)
	return c.SendStatus(200)
}

func UserRoutes(router fiber.Router) {
	user := router.Group("/users")
	user.Get("/", ListUsers)
	user.Post("/", CreateUser)
	user.Get("/:id", GetUser)
	user.Put("/:id", UpdateUser)
	user.Delete("/:id", DeleteUser)
}

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
	err = c.BodyParser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the user
	user.ID = uuid.New()
	// Create the User and return error if encountered
	_, err = user.Create(db)
	if err != nil {
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
	_, err = user.Get(db)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "The requested user could not be found", "data": err})
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) (err error) {
	db := database.DB
	user := models.User{}
	user.ID, err = uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "The requested user could not be found", "data": err})
	}
	_, err = user.Get(db)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "The requested user could not be found", "data": err})
	}
	err = c.BodyParser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	var u *models.User
	u, err = user.Update(db)

	return c.JSON(u)
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

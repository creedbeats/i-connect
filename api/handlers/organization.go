package handlers

import (
	"github.com/creedbeats/i-connect.git/api/database"
	"github.com/creedbeats/i-connect.git/api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ListOrganizations(c *fiber.Ctx) error {
	db := database.DB
	organization := models.Organization{}
	organizations, err := organization.List(db)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	return c.JSON(organizations)
}

func CreateOrganization(c *fiber.Ctx) (err error) {
	db := database.DB
	organization := models.Organization{}
	if err = c.BodyParser(&organization); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Create the Organization and return error if encountered
	if err = organization.Create(db); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create organization", "data": err})
	}
	return c.JSON(organization)
}

func GetOrganization(c *fiber.Ctx) (err error) {
	db := database.DB
	organization := models.Organization{}
	organization.ID, err = uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	if err = organization.Get(db); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "The requested organization could not be found", "data": err})
	}
	return c.JSON(organization)
}

func UpdateOrganization(c *fiber.Ctx) (err error) {
	db := database.DB
	organization := models.Organization{}
	
	if organization.ID, err = uuid.Parse(c.Params("id")); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "The requested organization could not be found", "data": err})
	}
	if err = c.BodyParser(&organization); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	if err = organization.Get(db); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "The requested organization could not be found", "data": err})
	}
	
	err = organization.Update(db)

	return c.JSON(organization)
}

func DeleteOrganization(c *fiber.Ctx) (err error) {
	db := database.DB
	organization := models.Organization{}
	organization.ID, err = uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	organization.Delete(db)
	return c.SendStatus(200)
}

func OrganizationRoutes(router fiber.Router) {
	organization := router.Group("/organizations")
	organization.Get("/", ListOrganizations)
	organization.Post("/", CreateOrganization)
	organization.Get("/:id", GetOrganization)
	organization.Put("/:id", UpdateOrganization)
	organization.Delete("/:id", DeleteOrganization)
}

package controllers

import "github.com/gofiber/fiber/v2"

func ListUsers(c *fiber.Ctx) error {
	return c.JSON("List Users")
}

func CreateUser(c *fiber.Ctx) error {
	return c.JSON("Create User")
}

func GetUser(c *fiber.Ctx) error {
	return c.JSON("Get User")
}

func UpdateUser(c *fiber.Ctx) error {
	return c.JSON("Update User")
}

func DeleteUser(c *fiber.Ctx) error {
	return c.JSON("Update User")
}

package controllers

import fiber "github.com/gofiber/fiber/v2"

func ListPosts(c *fiber.Ctx) error {
	return c.JSON("List Posts")
}

func CreatePost(c *fiber.Ctx) error {
	return c.JSON("Create Post")
}

func GetPost(c *fiber.Ctx) error {
	return c.JSON("Get Post")
}

func UpdatePost(c *fiber.Ctx) error {
	return c.JSON("Update Post")
}

func DeletePost(c *fiber.Ctx) error {
	return c.JSON("Update Post")
}

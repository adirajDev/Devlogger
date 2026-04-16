package handler

import "github.com/gofiber/fiber/v3"

func Test(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "Success",
		"message": "Hey! the test api is working!!!",
		"data":    nil,
	})
}

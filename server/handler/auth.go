package handler

import "github.com/gofiber/fiber/v3"

func Login(c fiber.Ctx) error {

	return c.JSON(fiber.Map{"status": "have fun"})
}

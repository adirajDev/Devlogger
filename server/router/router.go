package router

import (
	"github.com/adirajDev/Devlogger/server/handler"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api")
	api.Get("/", handler.Test)
}

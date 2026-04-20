package router

import (
	"github.com/adirajDev/Devlogger/server/handler"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api")
	api.Get("/", handler.Test)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/signup", handler.Signup)
	auth.Post("/login", handler.Login)
}

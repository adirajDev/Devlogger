package main

import (
	"log"

	"github.com/adirajDev/Devlogger/server/config"
	"github.com/adirajDev/Devlogger/server/database"
	"github.com/adirajDev/Devlogger/server/router"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	database.Connect()

	router.SetupRoutes(app)

	Port := config.GetEnv("PORT")
	log.Fatal(app.Listen(":" + Port))
}

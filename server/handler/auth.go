package handler

import (
	"context"

	"github.com/adirajDev/Devlogger/server/model"
	"github.com/adirajDev/Devlogger/server/util"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Signup(c fiber.Ctx) error {
	var input model.User
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on login request",
			"errors":  err.Error(),
		})
	}

	// Check if username/ email already exists
	if err := util.CheckIfUserNameExists(input.Username); err != nil {
		return err
	}

	if err := util.CheckIfUserEmailExists(input.Email); err != nil {
		return err
	}

	// Hash password
	hashedPassword, err := util.HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error hashing password",
		})
	}

	// Create user
	newUser := model.User{
		ID:        bson.NewObjectID(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Username:  input.Username,
		Password:  hashedPassword,
	}

	_, err = util.UserCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// TODO: set JWT cookie and return it

	return c.JSON(fiber.Map{"status": "have fun"})
}

func Login(c fiber.Ctx) error {

	return c.JSON(fiber.Map{"status": "have fun"})
}

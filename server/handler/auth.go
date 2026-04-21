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
	// Get username/Email
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}

	type UserData struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	input := new(LoginInput)
	var ud UserData
	if err := c.Bind().Body(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on login request",
			"errors":  err.Error(),
		})
	}

	//
	identity := input.Identity
	password := input.Password
	userModel, err := new(model.User), *new(error)

	if util.CheckEmailValidOrNot(identity) {
		userModel, err = util.GetUserByEmail(identity)
	} else {
		userModel, err = util.GetUserByUsername(identity)
	}

	const dummyHash = "$2a$10$7zFqzDbD3RrlkMTczbXG9OWZ0FLOXjIxXzSZ.QZxkVXjXcx7QZQiC" // to check pass hash when no user found

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal Server Error",
			"errors":  err,
		})
	} else if userModel == nil {
		// Always perform a hash check, even if the user doesn't exist, to prevent timing attacks
		util.CheckPasswordHash(password, dummyHash)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid identity or password",
			"errors":  err,
		})
	} else {
		ud = UserData{
			ID:       userModel.ID.Hex(),
			Username: userModel.Username,
			Email:    userModel.Email,
			Password: userModel.Password,
		}
	}

	if !util.CheckPasswordHash(password, ud.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid identity or password",
			"errors":  err,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success Login",
		"data": fiber.Map{
			"id":       ud.ID,
			"username": ud.Username,
			"email":    ud.Email,
		},
	})
}

package util

import (
	"context"
	"fmt"

	"github.com/adirajDev/Devlogger/server/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(mail string) (*model.User, error) {
	var user model.User
	err := UserCollection.FindOne(context.Background(), bson.M{"email": mail}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CheckIfUserNameExists(username string) error {
	err := UserCollection.FindOne(
		context.TODO(),
		bson.M{"username": username},
	)

	// TODO: handle error of database error
	if err == nil {
		return fmt.Errorf("Username already exists")
	}

	return nil
}

func CheckIfUserEmailExists(email string) error {
	err := UserCollection.FindOne(
		context.TODO(),
		bson.M{"email": email},
	)

	// TODO: handle error of database error
	if err == nil {
		return fmt.Errorf("Email already exists")
	}

	return nil
}

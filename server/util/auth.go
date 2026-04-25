package util

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/adirajDev/Devlogger/server/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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

func GetUserByEmail(mail string) (*model.User, error) {
	var user model.User
	err := UserCollection.FindOne(context.Background(), bson.M{"email": mail}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := UserCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CheckEmailValidOrNot(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func CheckIfUserExists(username string, email string) error {
	var result bson.M

	filter := bson.M{
		"$or": []bson.M{
			{"username": username},
			{"email": email},
		},
	}

	err := UserCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}

		return err
	}

	if result["username"] == username {
		return fmt.Errorf("username already exists")
	}
	if result["email"] == email {
		return fmt.Errorf("email already exists")
	}

	return fmt.Errorf("user already exists")
}

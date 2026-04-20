package util

import (
	"context"

	"github.com/adirajDev/Devlogger/server/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

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

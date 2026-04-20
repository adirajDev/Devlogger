package util

import "go.mongodb.org/mongo-driver/v2/mongo"

var UserCollection *mongo.Collection

func CreateDB(client *mongo.Client) {
	DB := client.Database("devlogger")

	UserCollection = DB.Collection("users")
}

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID         bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName  string        `json:"first_name" bson:"first_name"`
	LastName   string        `json:"last_name" bson:"last_name"`
	Email      string        `json:"email" bson:"email"`
	Username   string        `json:"username" bson:"username"`
	Password   string        `json:"password" bson:"password"`
	Updated_at time.Time     `json:"updated_at" bson:"updated_at"`
	Created_at time.Time     `json:"created_at" bson:"created_at"`
}

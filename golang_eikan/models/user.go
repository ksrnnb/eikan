package models

import (
	"context"
	"fmt"
	"log"

	"github.com/ksrnnb/eikan/db"
	"github.com/ksrnnb/eikan/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

// User model
// bsonも設定しておく。
// insertするときに内部でbsonにMarshalするので、そのときに空の場合は省略される。
type User struct {
	ID         string `json:"_id,omitempty" bson:"_id,omitempty"`
	Email      string `json:"email" bson:"email"`
	Password   string `json:"password" bson:"password"`
	Birthday   string `json:"birthday" bson:"birthday"` // string??
	GenderType string `json:"genderType" bson:"genderType"`
	CreatedAt  string `json:"createdAt" bson:"createdAt"`
	UpdatedAt  string `json:"updatedAt" bson:"updatedAt"`
	DeletedAt  string `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}

func collection() *mongo.Collection {
	name := "users"
	database, err := db.ConnectEikanDB()

	if err != nil {
		log.Printf("DB connection error: %v", err)
		return nil
	}

	return database.Collection(name)
}

// Create creates new user
func (user *User) Create() {
	user.CreatedAt = utils.CurrentTime()
	user.UpdatedAt = utils.CurrentTime()

	_, err := collection().InsertOne(context.Background(), user)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}
}

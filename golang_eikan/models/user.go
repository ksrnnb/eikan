package models

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ksrnnb/eikan/db"
	"github.com/ksrnnb/eikan/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

// User model
// bsonも設定しておく。
// insertするときに内部でbsonにMarshalするので、そのときに空の場合は省略される。
type User struct {
	ID         string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Email      string    `json:"email" bson:"email"`
	Password   string    `json:"password" bson:"password"`
	Birthday   string    `json:"birthday" bson:"birthday"` // string??
	GenderType int       `json:"genderType" bson:"genderType"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
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

// RegisterUser register user...
func RegisterUser(requestBody []byte) error {
	var user User
	err := json.Unmarshal(requestBody, &user)

	if err != nil {
		return err
	}

	// ユーザー作成
	err = user.Create()
	return err
}

// Create creates new user
func (user *User) Create() error {
	user.Password = utils.GenerateHashedPassword(user.Password)
	user.CreatedAt = utils.CurrentTime()
	user.UpdatedAt = utils.CurrentTime()

	_, err := collection().InsertOne(context.Background(), user)

	return err
}

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
type User struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Birthday   string `json:"birthday"` // string??
	GenderType string `json:"genderType"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	DeletedAt  string `json:"deletedAt"`
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

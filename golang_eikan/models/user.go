package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ksrnnb/eikan/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// User model
// bsonも設定しておく。
// insertするときに内部でbsonにMarshalするので、そのときに空の場合は省略される。
type User struct {
	ID         string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Email      string    `json:"email" bson:"email,omitempty"`
	Password   string    `json:"password" bson:"password,omitempty"`
	Birthday   string    `json:"birthday" bson:"birthday,omitempty"`
	GenderType int       `json:"genderType" bson:"genderType,omitempty"`
	UserToken  string    `json:"userToken" bson:"userToken,omitempty"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`
	DeletedAt  time.Time `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}

// RegisterUser register user...
func RegisterUser(requestBody []byte) error {
	var user User
	err := json.Unmarshal(requestBody, &user)

	if err != nil {
		return err
	}

	// ユーザー作成
	err = user.Register()
	return err
}

// Create メールアドレスとトークンの保存
func (user *User) Create() error {
	user.UserToken = utils.GenerateToken()
	user.CreatedAt = utils.CurrentTime()
	_, err := collection("users").InsertOne(context.Background(), user)

	return err
}

// Register creates new user
func (user *User) Register() error {
	user.Password = utils.GenerateHashedPassword(user.Password)
	user.UpdatedAt = utils.CurrentTime()

	filter := bson.D{
		{Key: "email", Value: user.Email},
		{Key: "userToken", Value: user.UserToken},
	}

	_, err := collection("users").UpdateOne(context.Background(), filter, user)

	return err
}

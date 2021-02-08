package models

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/ksrnnb/eikan/utils"
)

// Verify model
// bsonも設定しておく。
// insertするときに内部でbsonにMarshalするので、そのときに空の場合は省略される。
type Verify struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string    `json:"email" bson:"email"`
	Code      string    `json:"code" bson:"code"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`
}

// VerifyCode ...
func VerifyCode(requestBody []byte) (string, error) {
	var verify Verify
	err := json.Unmarshal(requestBody, &verify)

	if err != nil {
		return "", err
	}

	// 認証レコードを削除
	err = verify.Delete()
	if err != nil {
		utils.ErrorLog("ERROR", err)
		return "", err
	}

	// ユーザーを登録
	user := User{Email: verify.Email}
	err = user.Create()

	return user.UserToken, err
}

// RegisterEmail register email
func RegisterEmail(requestBody []byte) error {
	var verify Verify
	err := json.Unmarshal(requestBody, &verify)

	if err != nil {
		return err
	}

	// 認証レコード作成
	err = verify.Create()
	return err
}

// Create creates new verify code
func (verify *Verify) Create() error {
	verify.CreatedAt = utils.CurrentTime()
	verify.UpdatedAt = utils.CurrentTime()
	verify.Code = utils.GenerateVerifyCode()

	_, err := collection("verify_users").InsertOne(context.Background(), verify)

	return err
}

// Delete delete verify code
func (verify *Verify) Delete() error {

	result, err := collection("verify_users").DeleteOne(context.Background(), verify)

	if result.DeletedCount == 0 {
		return errors.New("verify record didn't exist")
	}

	return err
}

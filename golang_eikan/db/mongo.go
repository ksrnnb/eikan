package db

import (
	"context"
	"errors"
	"time"

	"github.com/ksrnnb/eikan/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client can access MongoDB

var endpoint string
var username string
var password string
var ctx context.Context
var client *mongo.Client

func init() {
	cfg := configs.New()
	endpoint = cfg["DB.ENDPOINT"]
	username = cfg["DB.USERNAME"]
	password = cfg["DB.PASSWORD"]
}

// Connect create connection to MongoDB and close connection end of this method.
func Connect() (*mongo.Client, error) {
	if client != nil {
		return client, nil
	}
	// https://docs.mongodb.com/drivers/go/
	// context.TODO()のほうがいい？
	// https://godoc.org/context#Background
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	// cancelは時間経過後にcloseさせる
	defer cancel()

	// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#example-Connect-SCRAM
	credential := getCredential()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(endpoint).SetAuth(credential))

	return client, err
}

// ConnectEikanDB returns eikan db connection
func ConnectEikanDB() (*mongo.Database, error) {
	if client == nil {
		return nil, errors.New("No client has created")
	}
	return client.Database("eikan"), nil
}

func getCredential() options.Credential {
	return options.Credential{
		Username: username,
		Password: password,
	}
}

// Close close database connection
// コネクションは切ったり接続したりしないのがベストプラクティス。
// https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
func Close() {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

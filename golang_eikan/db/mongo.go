package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ksrnnb/eikan/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var endpoint string
var username string
var password string
var client *mongo.Client
var ctx context.Context

func init() {
	cfg := configs.New()
	endpoint = cfg["DB.ENDPOINT"]
	username = cfg["DB.USERNAME"]
	password = cfg["DB.PASSWORD"]
}

// Connect create connection to MongoDB and close connection end of this method.
func Connect() {
	// https://docs.mongodb.com/drivers/go/
	// context.TODO()のほうがいい？
	// https://godoc.org/context#Background
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#example-Connect-SCRAM
	credential := options.Credential{
		Username: username,
		Password: password,
	}

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(endpoint).SetAuth(credential))
	if err != nil {
		panic(err)
	}

	// TODO: 閉じるタイミングを考える必要あり。
	defer Close()

	collection := client.Database("test").Collection("trainers")

	// You will be using this Trainer type later in the program
	// 大文字にしないとマーシャルされない、たぶん。
	type Trainer struct {
		Name string
		Age  int
		City string
	}

	ash := Trainer{"Ash", 10, "Pallet Town"}
	// misty := Trainer{"Misty", 10, "Cerulean City"}
	// brock := Trainer{"Brock", 15, "Pewter City"}

	tmp, err := collection.InsertOne(context.TODO(), ash)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("tmp: ", tmp)
	}
}

// Close close database connection
// コネクションは切ったり接続したりしないのがベストプラクティス。
// アプリケーション起動中に接続しっぱなしもおかしい？
// TODO: 閉じるタイミングを考える必要あり。
// https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
func Close() {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

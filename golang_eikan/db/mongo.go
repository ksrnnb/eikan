package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: ハードコード修正
func MongoInit() {
	// https://docs.mongodb.com/drivers/go/
	uri := "mongodb://mongo:27017/test"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#example-Connect-SCRAM
	credential := options.Credential{
		Username: "root",
		Password: "root",
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(credential))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("test").Collection("testCollection")

	type Test struct {
		name  string
		email string
	}
	test := Test{
		name:  "test",
		email: "test@test",
	}

	tmp, err := collection.InsertOne(ctx, test)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("tmp: ", tmp)
	}
	// Ping the primary
	// if err := client.Ping(ctx, readpref.Primary()); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Successfully connected and pinged.")
}

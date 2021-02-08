package models

import (
	"log"

	"github.com/ksrnnb/eikan/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func collection(collectionName string) *mongo.Collection {
	database, err := db.ConnectEikanDB()

	if err != nil {
		log.Printf("DB connection error: %v", err)
		return nil
	}

	return database.Collection(collectionName)
}

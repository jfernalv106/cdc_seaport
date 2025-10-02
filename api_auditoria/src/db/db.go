package db

import (
	"api_auditoria/src/config"
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(colleccion string, mongoDb config.Mongo) *mongo.Database {
	// Find .evn

	// Get value from .env
	MONGO_URI := fmt.Sprintf("mongodb://%s:%s@%s:%d", mongoDb.User, mongoDb.Password, mongoDb.Host, mongoDb.Port)

	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection.
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to mongoDB!!!")
	}

	fmt.Println("Connected to db")
	db := client.Database(mongoDb.Database)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func InitLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

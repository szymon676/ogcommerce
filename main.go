package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/szymon676/ogcommerce/api"
	"github.com/szymon676/ogcommerce/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		panic(err)
	}
	pstore := store.NewMongoProductStore(client.Database("mongo-products"))
	astore := store.NewMongoUsersStore(client.Database("mongo-users"))

	server := api.NewApiServer(os.Getenv("PORT"), pstore, astore)
	server.Run()
}

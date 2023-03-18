package main

import (
	"context"

	"github.com/szymon676/ogcommerce/api"
	"github.com/szymon676/ogcommerce/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	store := store.NewMongoProductStore(client.Database("mongo-products"))

	server := api.NewProductHandler(store, ":4000")
	server.Run()
}

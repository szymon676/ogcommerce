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
	pstore := store.NewMongoProductStore(client.Database("mongo-products"))
	astore := store.NewMongoUsersStore(client.Database("mongo-users"))

	server := api.NewApiServer(":4000", pstore, astore)
	server.Run()
}

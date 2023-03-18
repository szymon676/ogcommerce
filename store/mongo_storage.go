package store

import (
	"context"

	"github.com/szymon676/ogcommerce/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductStore struct {
	db   *mongo.Database
	coll string
}

func NewMongoProductStore(db *mongo.Database) *MongoProductStore {
	return &MongoProductStore{
		db:   db,
		coll: "products",
	}
}

func (s MongoProductStore) InsertProduct(ctx context.Context, product types.Product) error {
	product.ID = primitive.NewObjectID().Hex()
	_, err := s.db.Collection(s.coll).InsertOne(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (s MongoProductStore) GetProducts(ctx context.Context) ([]*types.Product, error) {
	cursor, err := s.db.Collection(s.coll).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	products := []*types.Product{}
	err = cursor.All(ctx, &products)
	return products, err
}

func (s MongoProductStore) GetProductByName(ctx context.Context, name string) (*types.Product, error) {
	var (
		res = s.db.Collection(s.coll).FindOne(ctx, bson.M{"name": name})
		p   = &types.Product{}
		err = res.Decode(p)
	)
	return p, err
}

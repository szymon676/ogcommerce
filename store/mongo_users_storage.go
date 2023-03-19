package store

import (
	"context"

	"github.com/szymon676/ogcommerce/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUsersStore struct {
	db   *mongo.Database
	coll string
}

func NewMongoUsersStore(db *mongo.Database) *MongoUsersStore {
	return &MongoUsersStore{
		db:   db,
		coll: "users",
	}
}

func (u MongoUsersStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	return users, nil
}

func (u MongoUsersStore) InsertUser(ctx context.Context, user types.User) error {
	user.ID = primitive.NewObjectID().Hex()
	_, err := u.db.Collection(u.coll).InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u MongoUsersStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var (
		res  = u.db.Collection(u.coll).FindOne(ctx, bson.M{"email": email})
		user = &types.User{}
		err  = res.Decode(user)
	)
	return user, err
}

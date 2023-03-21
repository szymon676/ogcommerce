package store

import (
	"context"

	"github.com/szymon676/ogcommerce/types"
)

type ProductsStorager interface {
	GetProducts(ctx context.Context) ([]*types.Product, error)
	InsertProduct(ctx context.Context, data types.Product) error
	GetProductByName(ctx context.Context, name string) (*types.Product, error)
	UpdateProductByName(ctx context.Context, name string, product types.Product) error
	DeleteProductByName(ctx context.Context, name string) error
}

type UsersStorager interface {
	GetUsers(ctx context.Context) ([]*types.User, error)
	InsertUser(ctx context.Context, user types.User) error
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
}

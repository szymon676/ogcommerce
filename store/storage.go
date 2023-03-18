package store

import (
	"context"

	"github.com/szymon676/ogcommerce/types"
)

type Storager interface {
	GetProducts(ctx context.Context) ([]*types.Product, error)
	InsertProduct(ctx context.Context, data types.Product) error
	GetProductByName(ctx context.Context, name string) (*types.Product, error)
}

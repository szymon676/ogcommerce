package types

import (
	"fmt"
)

type Product struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type ReqProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewProductFromRequest(reqp ReqProduct) (*Product, error) {
	if err := ValidateProduct(reqp); err != nil {
		return nil, err
	}

	return &Product{
		Name:        reqp.Name,
		Description: reqp.Description,
	}, nil
}

func ValidateProduct(rp ReqProduct) error {
	if len(rp.Description) < 5 {
		return fmt.Errorf("Product description must be at least 5 characters")
	}
	if len(rp.Name) < 3 {
		return fmt.Errorf("Product name must be at least 3 characters")
	}
	return nil
}

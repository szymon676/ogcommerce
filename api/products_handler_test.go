package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/szymon676/ogcommerce/store"
	"github.com/szymon676/ogcommerce/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestHandleGetProducts(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	pstore := store.NewMongoProductStore(client.Database("mongo-products"))
	p := NewProductHandler(pstore)

	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(makeHTTPHanler(p.handleGetProducts))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHandleCreateProduct(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	pstore := store.NewMongoProductStore(client.Database("mongo-products"))
	p := NewProductHandler(pstore)

	product := types.Product{
		Name:        "test product",
		Description: "test description",
	}

	jsonProduct, _ := json.Marshal(product)

	req, err := http.NewRequest("POST", "/products", bytes.NewReader(jsonProduct))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(makeHTTPHanler(p.handleGetProducts))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

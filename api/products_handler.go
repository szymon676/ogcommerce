package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/szymon676/ogcommerce/store"
	"github.com/szymon676/ogcommerce/types"
)

type ProductsHandler struct {
	store store.ProductsStorager
}

func NewProductHandler(store store.ProductsStorager) *ProductsHandler {
	return &ProductsHandler{
		store: store,
	}
}

func (h ProductsHandler) handleCreateProduct(w http.ResponseWriter, r *http.Request) error {
	var reqProduct types.ReqProduct

	if err := json.NewDecoder(r.Body).Decode(&reqProduct); err != nil {
		return err
	}
	product, err := types.NewProductFromRequest(reqProduct)
	if err != nil {
		return err
	}
	if err := h.store.InsertProduct(r.Context(), *product); err != nil {
		return err
	}

	return WriteJSON(w, 202, "Product created successfully")
}

func (h ProductsHandler) handleGetProducts(w http.ResponseWriter, r *http.Request) error {
	products, err := h.store.GetProducts(r.Context())
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, products)
}

func (h ProductsHandler) handleGetProduct(w http.ResponseWriter, r *http.Request) error {
	path := mux.Vars(r)
	product, err := h.store.GetProduct(r.Context(), path["id"])

	if err != nil {
		return err
	}

	return WriteJSON(w, 200, product)
}

func (h ProductsHandler) handleUpdateProduct(w http.ResponseWriter, r *http.Request) error {
	var reqProduct types.ReqProduct
	path := mux.Vars(r)

	if err := json.NewDecoder(r.Body).Decode(&reqProduct); err != nil {
		return err
	}

	product, err := types.NewProductFromRequest(reqProduct)
	if err != nil {
		return err
	}

	if err := h.store.UpdateProduct(r.Context(), path["id"], *product); err != nil {
		return err
	}

	return WriteJSON(w, 202, "updated product successfully")
}

func (h ProductsHandler) handleDeleteProduct(w http.ResponseWriter, r *http.Request) error {
	path := mux.Vars(r)

	if err := h.store.DeleteProduct(r.Context(), path["id"]); err != nil {
		return err
	}

	return WriteJSON(w, 202, "Product Deleted successfully")
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHanler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, statuscode int, body ...any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statuscode)

	return json.NewEncoder(w).Encode(body)
}

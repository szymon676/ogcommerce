package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/szymon676/ogcommerce/store"
	"github.com/szymon676/ogcommerce/types"
)

type ProductsHandler struct {
	store      store.Storager
	listenaddr string
}

func NewProductHandler(store store.Storager, listenaddr string) *ProductsHandler {
	return &ProductsHandler{
		listenaddr: listenaddr,
		store:      store,
	}
}

func (p ProductsHandler) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/products", makeHTTPHanler(p.handlePostProduct)).Methods("POST")
	router.HandleFunc("/products", makeHTTPHanler(p.handleGetProducts)).Methods("GET")
	router.HandleFunc("/products/{name}", makeHTTPHanler(p.handleGetProductByName)).Methods("GET")

	log.Println("server listening on port:", p.listenaddr)
	http.ListenAndServe(p.listenaddr, router)
}

func (h ProductsHandler) handlePostProduct(w http.ResponseWriter, r *http.Request) error {
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

func (h ProductsHandler) handleGetProductByName(w http.ResponseWriter, r *http.Request) error {
	path := mux.Vars(r)
	product, err := h.store.GetProductByName(r.Context(), path["name"])

	if err != nil {
		return err
	}

	return WriteJSON(w, 200, product)
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

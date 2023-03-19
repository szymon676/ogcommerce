package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/szymon676/ogcommerce/store"
)

type ApiServer struct {
	listenaddr string
	pstore     store.ProductsStorager
	astore     store.UsersStorager
}

func NewApiServer(addr string, pstore store.ProductsStorager, astore store.UsersStorager) *ApiServer {
	return &ApiServer{
		listenaddr: addr,
		pstore:     pstore,
		astore:     astore,
	}
}

func (s ApiServer) Run() {
	router := mux.NewRouter()
	p := NewProductHandler(s.pstore)
	ah := NewAuthHandler(s.astore)

	router.HandleFunc("/products", makeHTTPHanler(p.handlePostProduct)).Methods("POST")
	router.HandleFunc("/products", makeHTTPHanler(p.handleGetProducts)).Methods("GET")
	router.HandleFunc("/products/{name}", makeHTTPHanler(p.handleGetProductByName)).Methods("GET")
	router.HandleFunc("/products/{name}", makeHTTPHanler(p.handleDeleteProductByName)).Methods("DELETE")

	router.HandleFunc("/login", makeHTTPHanler(ah.handleLoginUser)).Methods("POST")

	log.Println("server listening on port:", s.listenaddr)
	http.ListenAndServe(s.listenaddr, router)
}

package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"mymodule8-rest-with-gorilla/data"
	"mymodule8-rest-with-gorilla/middleware"
	"net/http"
	"strconv"
)

type ProductsHandler struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{
		l: l,
	}
}

func (self *ProductsHandler) GetProducts(rw http.ResponseWriter, req *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to json encode", http.StatusInternalServerError)
	}
}

func (self *ProductsHandler) AddProduct(rw http.ResponseWriter, req *http.Request) {
	self.l.Println("Handle POST /products")
	product := req.Context().Value(middleware.KeyProduct{}).(*data.Product)

	// ---- add product to storage
	data.AddProduct(product)
	rw.WriteHeader(http.StatusCreated)
}

func (self *ProductsHandler) UpdateProduct(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "atoi fail", http.StatusBadRequest)
	}

	product := req.Context().Value(middleware.KeyProduct{}).(*data.Product)

	// ---- add product to storage
	err = data.UpdateProduct(id, product)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.WriteHeader(http.StatusOK)
}

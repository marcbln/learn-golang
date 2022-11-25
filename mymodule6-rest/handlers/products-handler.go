package handlers

import (
	"log"
	"mymodule6-rest/data"
	"net/http"
)

type ProductsHandler struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{
		l: l,
	}
}

func (self *ProductsHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	// ---- handle get
	if req.Method == http.MethodGet {
		self.getProducts(rw, req)
		return
	}

	// ---- handle update (PUT)
	if req.Method == http.MethodPut {
		http.Error(rw, "TODO: implement PUT", http.StatusNotImplemented)
		return
	}

	// catch-all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (self *ProductsHandler) getProducts(rw http.ResponseWriter, req *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to json encode", http.StatusInternalServerError)
	}
}

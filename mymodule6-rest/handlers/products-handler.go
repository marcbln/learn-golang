package handlers

import (
	"encoding/json"
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
	products := data.GetProducts()
	ret, err := json.Marshal(products)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	rw.Write(ret)
}

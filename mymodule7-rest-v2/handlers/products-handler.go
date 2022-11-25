package handlers

import (
	"log"
	"mymodule7-rest-v2/data"
	"net/http"
	"regexp"
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

func (self *ProductsHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	// ---- handle get
	if req.Method == http.MethodGet {
		self.getProducts(rw, req)
		return
	}

	// ---- handle create (POST)
	if req.Method == http.MethodPost {
		self.addProduct(rw, req)
		return
	}

	// ---- handle update (PUT)
	if req.Method == http.MethodPut {
		self.l.Println("PUT", req.URL.Path)
		// expect the id in the URI
		regex := regexp.MustCompile(`/([0-9]+)`)
		path := req.URL.Path // products/123
		gr := regex.FindAllStringSubmatch(path, -1)
		self.l.Printf("%#v", gr)
		if len(gr[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}

		id, err := strconv.Atoi(gr[0][1])
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}
		// self.l.Printf("id: %v", id)
		self.updateProduct(id, rw, req)

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

func (self *ProductsHandler) addProduct(rw http.ResponseWriter, req *http.Request) {
	self.l.Println("Handle POST /products")
	// ---- deserialize product from json
	product := &data.Product{}
	err := product.FromJSON(req.Body)
	if err != nil {
		http.Error(rw, "json decoding failed", http.StatusBadRequest)
		return
	}
	self.l.Printf("%#v", product)

	// ---- add product to storage
	data.AddProduct(product)
	rw.WriteHeader(http.StatusCreated)
}

func (self *ProductsHandler) updateProduct(id int, rw http.ResponseWriter, req *http.Request) {
	// ---- deserialize product from json
	product := &data.Product{}
	err := product.FromJSON(req.Body)
	if err != nil {
		http.Error(rw, "json decoding failed", http.StatusBadRequest)
		return
	}
	self.l.Printf("%#v", product)

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

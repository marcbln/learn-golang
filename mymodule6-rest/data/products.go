package data

import "time"

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

var now = time.Now().UTC().String()

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Ice Tea",
		Description: "Tea with ice cubes",
		Price:       2.95,
		SKU:         "P-IceTea",
		CreatedAt:   now,
		UpdatedAt:   now,
	},
	{
		ID:          2,
		Name:        "Bubble Tea",
		Description: "Tea with bubbles",
		Price:       3.95,
		SKU:         "P-BuTea",
		CreatedAt:   now,
		UpdatedAt:   now,
	},
}

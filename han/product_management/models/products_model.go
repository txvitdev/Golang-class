package models

type Product struct {
	ID int
	Name string
	Price float32
}

type CreateProduct struct {
	Name string
	Price float32
}

type UpdateProduct struct {
	ID int
	Name string
	Price float32
}

type ProductResponse struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
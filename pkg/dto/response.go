package dto

type ProductResponse struct {
	Status  bool    `json:"status"`
	Product Product `json:"product"`
}

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
}

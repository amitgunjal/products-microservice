package models

type Product struct {
	ID       string  `json:"id" gorm:"PRIMARY_KEY"`
	Name     string  `json:"first_name" gorm:"not null"`
	Price    float64 `json:"price" gorm:"not null"`
	Quantity int64   `json:"quantity" gorm:"not null"`
}

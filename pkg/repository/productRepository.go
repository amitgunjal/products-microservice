package repository

import "github.com/jinzhu/gorm"

type ProductRepository interface {
}
type ProductRepositoryImpl struct {
	Database *gorm.DB
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

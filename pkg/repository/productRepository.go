package repository

import (
	"github.com/jinzhu/gorm"
	"products/pkg/models"
)

type ProductRepository interface {
	AddProduct(product models.Product) (*models.Product, error)
}
type ProductRepositoryImpl struct {
	Database *gorm.DB
}

func (p ProductRepositoryImpl) AddProduct(product models.Product) (*models.Product, error) {
	db := p.Database.Create(&product)
	if db.Error != nil {
		return nil, db.Error
	}
	return &product, nil
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

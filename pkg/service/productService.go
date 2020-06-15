package service

import "products/pkg/repository"

type ProductService interface {
}

type ProductServiceImpl struct {
	Repository repository.ProductRepository
}

func NewProductService() ProductService {
	return &ProductServiceImpl{}
}

package service

import (
	uuid "github.com/amitgunjal/unik"
	"github.com/jinzhu/copier"
	"products/pkg/dto"
	"products/pkg/models"
	"products/pkg/repository"
)

type ProductService interface {
	AddProduct(productRequest dto.ProductRequest) (*dto.ProductResponse, error)
}

type ProductServiceImpl struct {
	Repository repository.ProductRepository
}

func (p ProductServiceImpl) AddProduct(productRequest dto.ProductRequest) (*dto.ProductResponse, error) {
	var product models.Product
	err := copier.Copy(&product, productRequest)
	if err != nil {
		return nil, err
	}
	product.ID = uuid.NewUUID()
	res, err := p.Repository.AddProduct(product)
	if err != nil {
		return nil, err
	}
	var productResponse dto.Product
	err = copier.Copy(&productResponse, res)
	if err != nil {
		return nil, err
	}
	response := dto.ProductResponse{
		Status:  true,
		Product: productResponse,
	}
	return &response, nil
}

func NewProductService() ProductService {
	return &ProductServiceImpl{}
}

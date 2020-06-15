package container

import (
	"github.com/jinzhu/gorm"
	"log"
	"products/pkg/controllers"
	"products/pkg/repository"
	"products/pkg/service"
	"products/utils"
)

type Container struct {
	DB                    *gorm.DB
	ProductController     *controllers.ProductController
	ProductServiceImpl    *service.ProductServiceImpl
	ProductRepositoryImpl *repository.ProductRepositoryImpl
}

func (c *Container) GetProductController() *controllers.ProductController {
	if c.ProductController == nil {
		c.ProductController = &controllers.ProductController{
			Service: c.GetProductServiceImpl(),
		}
	}
	return c.ProductController
}

func (c *Container) GetProductServiceImpl() *service.ProductServiceImpl {
	if c.ProductServiceImpl == nil {
		c.ProductServiceImpl = &service.ProductServiceImpl{
			Repository: c.GetProductRepositoryImpl(),
		}
	}
	return c.ProductServiceImpl
}

func (c *Container) GetProductRepositoryImpl() *repository.ProductRepositoryImpl {
	if c.ProductRepositoryImpl == nil {
		c.ProductRepositoryImpl = &repository.ProductRepositoryImpl{
			Database: c.GetDb(),
		}
	}
	return c.ProductRepositoryImpl
}

func (c *Container) GetDb() *gorm.DB {
	var err error
	if c.DB != nil {
		return c.DB
	}
	c.DB, err = utils.StartUp()
	if err != nil {
		log.Fatal("DB connection failed", err)
	}
	return c.DB
}

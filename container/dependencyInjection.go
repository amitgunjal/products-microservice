package container

import (
	"github.com/jinzhu/gorm"
	"log"
	"products/pkg/controllers"
	"products/utils"
)

type Container struct {
	DB                *gorm.DB
	ProductController *controllers.ProductController
}

func (c *Container) GetProductController() *controllers.ProductController {
	if c.ProductController == nil {
		c.ProductController = &controllers.ProductController{
			//Service: c.GetUserServiceImpl(),
		}
	}
	return c.ProductController
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

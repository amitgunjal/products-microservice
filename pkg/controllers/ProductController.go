package controllers

import (
	"net/http"
	"products/pkg/service"
	"products/utils"
)

type ProductController struct {
	Service *service.ProductServiceImpl
}

func (c *ProductController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "I am healthy!")
	utils.Respond(w, resp)
}

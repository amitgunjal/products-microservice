package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"products/pkg/dto"
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

func (c *ProductController) AddProduct(w http.ResponseWriter, r *http.Request) {
	var request dto.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}
	log.Println(request)
	res, err := c.Service.AddProduct(request)

	utils.RespondJSON(w, http.StatusOK, res)
}

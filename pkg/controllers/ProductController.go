package controllers

import (
	"net/http"
	"products/utils"
)

type ProductController struct {

}
func (c *ProductController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "I am healthy!")
	utils.Respond(w, resp)
}

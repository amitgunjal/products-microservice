package router

import (
	"products/container"

	"github.com/gorilla/mux"
)

func InitRoutes(container *container.Container) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router = ProductRouters(container, router)
	return router
}

package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"products/container"
)

func ProductRouters(container *container.Container, router *mux.Router) *mux.Router {
	// Checking health of REST Endpoints
	router.Handle("/healthcheck", http.HandlerFunc(container.GetProductController().HealthCheck)).Methods(http.MethodGet)
	router.Handle("/product", http.HandlerFunc(container.GetProductController().AddProduct)).Methods(http.MethodPost)
	return router
}

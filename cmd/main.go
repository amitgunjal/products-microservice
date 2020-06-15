package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"products/constants"
	"products/container"
	"products/router"
	"products/utils"
)

func main() {
	controller := container.Container{}
	routes := router.InitRoutes(&controller)
	port := utils.GetServicePort()
	if port == "" {
		port = "9000" //localhost
	}

	fmt.Printf("Serving API's at URL: http://localhost:%s", port)
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{constants.AllowedHeaders}),
		handlers.AllowedOrigins([]string{constants.AllowedOrigins}),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPut}),
	)
	routes.Use(cors)
	err := http.ListenAndServe(":"+port, logRequest(routes)) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Println(err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

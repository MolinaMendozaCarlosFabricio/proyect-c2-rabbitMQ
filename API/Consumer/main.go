package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	products_routes "products_api.com/p/src/Products/infrastructure/routes"
	requests_routes "products_api.com/p/src/Requests/infrastructure/routes"
)

func main() {
	r := gin.Default()
	c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:4200"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
        AllowCredentials: true,
        Debug:            true,
    })
	products_routes.ProductRoutes(r)
	requests_routes.RequestRoutes(r)

	handler := c.Handler(r)

	http.ListenAndServe(":9080", handler)
}
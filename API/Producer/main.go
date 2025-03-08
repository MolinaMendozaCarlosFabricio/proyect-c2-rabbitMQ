package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	request_routes "request_api.com/r/src/requests/infrastructure/routes"
	user_routes "request_api.com/r/src/users/infrastructure/routes"
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

	user_routes.UserRoutes(r)
	request_routes.RequestRoutes(r)

	handler := c.Handler(r)

	http.ListenAndServe(":8080", handler)
}
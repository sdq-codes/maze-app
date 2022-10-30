package server

import (
	"github.com/gorilla/handlers"
	"github.com/sdq-codes/maze-app/router"
	"net/http"
)

// Server configures and returns a new http.Server
func Server() *http.Server {
	r := router.AllRoutes()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	srv := &http.Server{Handler: handlers.CORS(header, methods, origins)(r), Addr: ":9000"}
	return srv
}

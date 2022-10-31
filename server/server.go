package server

import (
	"github.com/gorilla/handlers"
	"github.com/sdq-codes/maze-app/router"
	"net/http"
	"os"
)

// Server configures and returns a new http.Server
func Server() *http.Server {
	r := router.AllRoutes()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	port := os.Getenv("PORT")
	if port == "" {
		port = "9099"
	}
	srv := &http.Server{Handler: handlers.CORS(header, methods, origins)(r), Addr: ":" + port}
	return srv
}

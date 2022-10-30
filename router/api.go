package router

import (
	"github.com/gorilla/mux"
	"github.com/sdq-codes/maze-app/app"
	"log"
	"net/http"
)

var (
	maze app.Maze
)

func AllRoutes() *mux.Router {
	R := mux.NewRouter()
	R.Handle("/", handlers.Handler(healthHandler.CheckHealth)).Methods(http.MethodGet)
	v1 := R.PathPrefix("/v1").Subrouter()
	v1.Handle("/mazes", handlers.Handler(maze.Solve)).Methods(http.MethodPost)
	return R

}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package router

import (
	"github.com/gorilla/mux"
	app "github.com/sdq-codes/maze-app/app/maze"
	"github.com/sdq-codes/maze-app/server/handlers"
	"log"
	"net/http"
)

var (
	mazeApp app.Maze
)

func AllRoutes() *mux.Router {
	R := mux.NewRouter()
	v1 := R.PathPrefix("/v1").Subrouter()
	v1.Handle("/maze/solve", handlers.Handler(mazeApp.Solve)).Methods(http.MethodPost)
	v1.Handle("/maze/generate", handlers.Handler(mazeApp.Create)).Methods(http.MethodGet)
	return R

}

func init() {
	mazeApp = app.NewMaze()
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

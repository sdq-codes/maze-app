package app

import (
	"github.com/sdq-codes/maze-app/handlers"
	"github.com/sdq-codes/maze-app/models"
	"github.com/sdq-codes/maze-app/server/decoder"
	"github.com/sdq-codes/maze-app/server/responses"
	"net/http"
)

type Maze interface {
	Solve(w http.ResponseWriter, r *http.Request) error
}

func NewMaze() Maze {
	return maze{}
}

type maze struct{}

func (u user) Solve(w http.ResponseWriter, r *http.Request) error {
	requestBody := &models.Maze{}
	if err := decoder.DecodeJSON(r.Body, requestBody); err != nil {
		return err
	}
	allRoutes := handlers.Maze.Solve(requestBody.Maze)
	return responses.Created(
		"Authentication successful",
		"User registration successful",
		responseBody,
	).ToJSON(w)
}

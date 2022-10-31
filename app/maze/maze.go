package app

import (
	"github.com/sdq-codes/maze-app/handlers"
	"github.com/sdq-codes/maze-app/models"
	"github.com/sdq-codes/maze-app/server/decoder"
	"github.com/sdq-codes/maze-app/server/responses"
	"net/http"
	"strconv"
)

type Maze interface {
	Solve(w http.ResponseWriter, r *http.Request) error
	Create(w http.ResponseWriter, r *http.Request) error
}

func NewMaze() Maze {
	return maze{
		handlers.Maze{},
	}
}

type maze struct {
	mazeHandler handlers.Maze
}

func (m maze) Solve(w http.ResponseWriter, r *http.Request) error {
	requestBody := &models.Maze{}
	if err := decoder.DecodeJSON(r.Body, requestBody); err != nil {
		return err
	}
	allRoutes := m.mazeHandler.Solve(requestBody.Maze, []string{})
	if len(allRoutes) == 0 {
		return responses.Created(
			"Maze solver failed",
			"No exit route found",
			"Sorry",
		).ToJSON(w)
	}
	return responses.Created(
		"Maze solver successful",
		"All exit route found",
		allRoutes,
	).ToJSON(w)
}

func (m maze) Create(w http.ResponseWriter, r *http.Request) error {
	maxNode, err := strconv.Atoi(r.URL.Query().Get("max-children"))
	if err != nil {
		return responses.Fail(
			"Maze creator failed",
			"Unable to fetch maximum allowed node",
			http.StatusBadRequest,
		).ToJSON(w)
	}
	mazeJSON := m.mazeHandler.CreateMaze(0, maxNode)
	return responses.Created(
		"Maze solver successful",
		"All exit route found",
		mazeJSON,
	).ToJSON(w)
}

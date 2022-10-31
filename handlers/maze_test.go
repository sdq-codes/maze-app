package handlers

import (
	"testing"
)

func TestMaze_Exit(t *testing.T) {
	mazeCase := map[string]interface{}{
		"forward": "exit",
	}
	maze := Maze{}
	response := maze.Solve(mazeCase, []string{})
	testResult := [][]string{
		[]string{
			"forward",
		},
	}
	for i := range testResult {
		for x := range testResult[i] {
			if testResult[i][x] != response[i][x] {
				t.Errorf("got %q, wanted %q", response, testResult)
			}
		}
	}
}

func TestMaze_Exit_Fail(t *testing.T) {
	mazeCase := map[string]interface{}{
		"forward": "ogre",
	}
	maze := Maze{}
	response := maze.Solve(mazeCase, []string{})
	testResult := [][]string{
		[]string{
			"forward",
		},
	}
	if len(response) > 0 {
		t.Errorf("got %q, wanted %q", testResult, response)
	}
}

func TestMaze_Solve_No_Exit(t *testing.T) {
	mazeCase := map[string]interface{}{
		"forward": "tiger",
		"left":    "ogre",
		"right":   "demon",
	}
	maze := Maze{}
	response := maze.Solve(mazeCase, []string{})
	testResult := [][]string{
		[]string{},
	}
	for i := range testResult {
		for x := range testResult[i] {
			if testResult[i][x] != response[i][x] {
				t.Errorf("got %q, wanted %q", response, testResult)
			}
		}
	}
}

func TestMaze_Solve_One_Exits(t *testing.T) {
	mazeCase := map[string]interface{}{
		"forward": "tiger",
		"left": map[string]interface{}{
			"forward": map[string]interface{}{
				"upstairs": "exit",
			},
			"left": "dragon",
		},
		"right": map[string]interface{}{
			"forward": "dead end",
		},
	}
	maze := Maze{}
	response := maze.Solve(mazeCase, []string{})
	testResult := [][]string{
		[]string{
			"left",
			"forward",
			"upstairs",
		},
	}
	for i := range testResult {
		for x := range testResult[i] {
			if testResult[i][x] != response[i][x] {
				t.Errorf("got %q, wanted %q", response, testResult)
			}
		}
	}
}

func TestMaze_Solve_Two_Exits(t *testing.T) {
	mazeCase := map[string]interface{}{
		"forward": "tiger",
		"left": map[string]interface{}{
			"forward": map[string]interface{}{
				"upstairs": "exit",
			},
			"left": "exit",
		},
		"right": map[string]interface{}{
			"forward": "dead end",
		},
	}
	maze := Maze{}
	response := maze.Solve(mazeCase, []string{})
	testResult := [][]string{
		[]string{
			"left",
			"left",
		},
		[]string{
			"left",
			"forward",
			"upstairs",
		},
	}
	for i := range testResult {
		for x := range testResult[i] {
			if testResult[i][x] != response[i][x] {
				t.Errorf("got %q, wanted %q", response, testResult)
			}
		}
	}
}

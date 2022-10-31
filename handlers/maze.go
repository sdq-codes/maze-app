package handlers

import (
	"github.com/sdq-codes/maze-app/enums"
	"math/rand"
	"sort"
)

type Maze struct {
	AllRoutes [][]string
}

func (m *Maze) CreateMaze(currentNodeLevel, maxNodeLevel int) map[string]interface{} {
	allowedDirection := enums.AllPossibleDirections
	maze := map[string]interface{}{}
	maxAllowedKeys := len(enums.AllPossibleDirections)
	for i := 1; i <= maxAllowedKeys; i++ {
		key := allowedDirection[rand.Intn(len(allowedDirection))]
		if currentNodeLevel+1 == maxNodeLevel {
			maze[key] = enums.AllPossibleValues[rand.Intn(len(enums.AllPossibleValues))]
			return maze
		}
		if rand.Intn(2)%2 == 0 {
			maze[key] = m.CreateMaze(currentNodeLevel+1, maxNodeLevel)
		} else {
			maze[key] = enums.AllPossibleValues[rand.Intn(len(enums.AllPossibleValues))]
		}
	}
	return maze
}

func (m *Maze) Solve(
	mazeItem interface{},
	keys []string,
) [][]string {
	mazeItemMap, okMap := mazeItem.(map[string]interface{})
	mazeItemString, okString := mazeItem.(string)
	if okMap && !okString {
		for key, element := range mazeItemMap {
			keys = append(keys, key)
			_ = m.Solve(element, keys)
			if len(keys) > 0 {
				keys = keys[:len(keys)-1]
			}
		}
	}
	if !okMap && okString {
		if mazeItemString == "exit" {
			m.AllRoutes = append(m.AllRoutes, keys)
			keys = keys[:len(keys)-1]
		}
	}
	if len(m.AllRoutes) > 1 {
		sort.Slice(m.AllRoutes, func(i, j int) bool {
			return len(m.AllRoutes[i]) < len(m.AllRoutes[j])
		})
	}
	return m.AllRoutes
}

//Solve for case where its a 1 dimensional array with only one possible answer
func (m *Maze) SolveBaseCaseWithOneExit(mazeItem map[string]string) [][]string {
	for key, element := range mazeItem {
		if element == "exit" {
			return [][]string{
				[]string{
					key,
				},
			}
		}
	}
	return [][]string{}
}

//Solve for case where its a 1 dimensional array with multiple possible answers
func (m *Maze) SolveBaseCaseWithMultipleAnswers(mazeItem map[string]string) [][]string {
	for key, element := range mazeItem {
		if element == "exit" {
			m.AllRoutes = append(m.AllRoutes, []string{
				key,
			})
		}
	}
	return m.AllRoutes
}

//Solve for case where its multiple  dimensional array with one possible answers
func (m *Maze) Solve2DArrayWithOnlyOneAnswer(
	mazeItem interface{},
	keys []string,
) [][]string {
	mazeItemMap, okMap := mazeItem.(map[string]interface{})
	mazeItemString, okString := mazeItem.(string)
	if okMap && !okString {
		for key, element := range mazeItemMap {
			keys = append(keys, key)
			_ = m.Solve(element, keys)
			if len(keys) > 0 {
				keys = keys[:len(keys)-1]
			}
		}
	}
	if !okMap && okString {
		if mazeItemString == "exit" {
			m.AllRoutes = append(m.AllRoutes, keys)
			sort.Slice(m.AllRoutes, func(i, j int) bool {
				return len(m.AllRoutes[i]) < len(m.AllRoutes[j])
			})
			return m.AllRoutes
		}
	}
	return m.AllRoutes
}

//Solve for case where its multiple  dimensional array with multiple possible answers
func (m *Maze) Solve2DArrayWithMultiplePossibleAnswers(
	mazeItem interface{},
	keys []string,
) [][]string {
	mazeItemMap, okMap := mazeItem.(map[string]interface{})
	mazeItemString, okString := mazeItem.(string)
	if okMap && !okString {
		for key, element := range mazeItemMap {
			keys = append(keys, key)
			_ = m.Solve(element, keys)
			if len(keys) > 0 {
				keys = keys[:len(keys)-1]
			}
		}
	}
	if !okMap && okString {
		if mazeItemString == "exit" {
			m.AllRoutes = append(m.AllRoutes, keys)
			keys = keys[:len(keys)-1]
		}
	}
	if len(m.AllRoutes) > 1 {
		sort.Slice(m.AllRoutes, func(i, j int) bool {
			return len(m.AllRoutes[i]) < len(m.AllRoutes[j])
		})
	}
	return m.AllRoutes
}

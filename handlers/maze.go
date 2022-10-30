package handlers

type Maze struct {
	AllRoutes [][]string
}

func (m *Maze) Solve(
	mazeItem interface{},
	keys []string,
) [][]string {
	//fmt.Println("outside dict", mazeItem, keys)
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
		}
	}
	return m.AllRoutes
}

func case1(mazeItemMap map[string]string) []string {
	for key, element := range mazeItemMap {
		if element == "exit" {
			return []string{key}
		}
	}
	return []string{}
}

func case2(mazeItemMap map[string]string) []string {
	var AllExitRoutes []string
	for key, element := range mazeItemMap {
		if element == "exit" {
			AllExitRoutes = append(AllExitRoutes, key)
		}
	}
	return AllExitRoutes
}

package solvekit

import parsekit "toolKit/parseKit"

func BFS(graph map[*parsekit.Room][]*parsekit.Room, start, end *parsekit.Room) []*parsekit.Room {
	if _, exists := graph[start]; !exists {
		return nil
	}
	if _, exists := graph[end]; !exists {
		return nil
	}

	dist := make(map[*parsekit.Room]int)
	parent := make(map[*parsekit.Room]*parsekit.Room)
	queue := []*parsekit.Room{start}
	dist[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			break
		}

		for _, neighbor := range graph[current] {
			if _, visited := dist[neighbor]; !visited {
				dist[neighbor] = dist[current] + 1
				parent[neighbor] = current
				queue = append(queue, neighbor)
			}
		}
	}

	if _, ok := dist[end]; !ok {
		return nil
	}

	path := []*parsekit.Room{}
	current := end
	for current != nil {
		path = append(path, current)
		current = parent[current]
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

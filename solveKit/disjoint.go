package solvekit

import parsekit "toolKit/parseKit"

func FindDisjointPaths(start, end *parsekit.Room) [][]*parsekit.Room {
	if start == nil || end == nil {
		return nil
	}

	var allPaths [][]*parsekit.Room

	tempGraph := make(map[*parsekit.Room][]*parsekit.Room)
	for _, roomPtr := range parsekit.Rooms {
		neighbors := make([]*parsekit.Room, len(roomPtr.Link))
		copy(neighbors, roomPtr.Link)
		tempGraph[roomPtr] = neighbors
	}

	for {
		path := BFS(tempGraph, start, end)
		if path == nil {
			break
		}

		allPaths = append(allPaths, path)

		for i := 1; i < len(path)-1; i++ {
			node := path[i]
			delete(tempGraph, node)

			for _, neighbors := range tempGraph {
				for j := 0; j < len(neighbors); j++ {
					if neighbors[j] == node {
						neighbors = append(neighbors[:j], neighbors[j+1:]...)
						j--
					}
				}
			}
		}
	}

	return allPaths
}

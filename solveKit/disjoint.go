package solvekit

import parsekit "toolKit/parseKit"

// FindDisjointPaths returns a slice of node-disjoint paths from start to end.
// Each path is a slice of *parsekit.Room, from start to end.
// Paths are found in order of increasing length (shortest first).
func FindDisjointPaths(start, end *parsekit.Room) [][]*parsekit.Room {
	if start == nil || end == nil {
		return nil
	}

	var allPaths [][]*parsekit.Room

	// Make a deep copy of the graph by cloning the Link slices
	// We'll work on a temporary graph so we don't destroy the original
	tempGraph := make(map[*parsekit.Room][]*parsekit.Room)
	for _, roomPtr := range parsekit.Rooms {
		// roomPtr is *parsekit.Room
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

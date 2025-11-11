package functions

import (
	"sort"
)

// CalculateTurns figures out how many ants go per path and the total number of turns needed.
func CalculateTurns(paths []Path, antNumber int) ([]int, int) {
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	assigned := AssignAnts(paths, antNumber)

	maxTurn := 0
	for i := range paths {
		turn := len(paths[i]) - 1 + assigned[i]
		if turn > maxTurn {
			maxTurn = turn
		}
	}

	return assigned, maxTurn
}

// AssignAnts distributes ants across paths to balance travel time.
func AssignAnts(paths []Path, antNumber int) []int {
	pathLen := make([]int, len(paths))

	for i, path := range paths {
		pathLen[i] = len(path)
	}

	assigned := make([]int, len(paths))
	antsLeft := antNumber

	for antsLeft > 0 {
		target := FindMinLoadPath(pathLen, assigned)
		assigned[target]++
		antsLeft--
	}

	return assigned
}

// FindMinLoadPath picks the path with the smallest combined length and assigned ants.
func FindMinLoadPath(pathLen, assigned []int) int {
	target := 0
	lowest := pathLen[0] + assigned[0]

	for i := 1; i < len(pathLen); i++ {
		load := pathLen[i] + assigned[i]
		if load <= lowest {
			target = i
			lowest = load
		}
	}

	return target
}

// Add inserts a node into the priority queue, keeping it sorted by priority.
func (queue *queue) Add(room Node) {
	*queue = append(*queue, room)
	sort.Slice(*queue, func(i, j int) bool {
		return (*queue)[i].Priority < (*queue)[j].Priority
	})
}

// Poll removes and returns the node with the smallest priority value.
func (queue *queue) Poll() Node {
	room := (*queue)[0]
	*queue = (*queue)[1:]
	return room
}

// buildPathfromParent reconstructs a path from the parent map between start and end.
func buildPathfromParent(parent map[string]string, start, end string) Path {
	path := Path{}
	for current := end; current != ""; current = parent[current] {
		path = append(Path{current}, path...)
		if current == start {
			break
		}
	}
	return path
}

// HasDuplicateRoomAcrossPaths checks if any room appears in more than one path.
func HasDuplicateRoomAcrossPaths(paths []Path) bool {
	seen := make(map[string]int)
	for i, path := range paths {
		for _, room := range path[:len(path)-1] {
			if prev, exists := seen[room]; exists && prev != i {
				return true
			}
			seen[room] = i
		}
	}
	return false
}

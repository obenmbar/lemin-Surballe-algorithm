package solvekit

import (
	"fmt"
	"strings"

	// adjust import path if needed
	parsekit "toolKit/parseKit"
)

type Ant struct {
	ID       int
	Path     []*parsekit.Room // full path: [start, ..., end]
	Step     int              // current index in path (0 = at start, not moved yet)
	Finished bool
}

func Simulate(antCount int, start, end *parsekit.Room) {
	// 1. Find disjoint paths
	paths := FindDisjointPaths(start, end)
	if len(paths) == 0 {
		return // no solution
	}

	// 2. Create ants and assign to paths (round-robin)
	ants := make([]Ant, antCount)
	for i := 0; i < antCount; i++ {
		ants[i] = Ant{
			ID:   i + 1,
			Path: paths[i%len(paths)],
			Step: 0,
		}
	}

	// 3. Simulate turns
	turn := 0
	for {
		turn++
		var moves []string
		occupied := make(map[string]bool) // room name → occupied this turn

		// Allow multiple ants to finish in same turn → don't block 'end'
		// But block all other rooms

		// First: plan moves (don't move yet)
		for i := range ants {
			ant := &ants[i]
			if ant.Finished {
				continue
			}

			// If ant is at end
			if ant.Step == len(ant.Path)-1 {
				ant.Finished = true
				continue
			}

			nextRoom := ant.Path[ant.Step+1]
			nextName := nextRoom.Name

			// Can move if:
			// - next room is END, OR
			// - next room is not occupied this turn
			if nextName == end.Name || !occupied[nextName] {
				occupied[nextName] = true
				ant.Step++
				moves = append(moves, fmt.Sprintf("L%d-%s", ant.ID, nextName))
			}
		}

		// Check if all finished
		allDone := true
		for _, ant := range ants {
			if !ant.Finished {
				allDone = false
				break
			}
		}

		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}

		if allDone {
			break
		}
	}
}

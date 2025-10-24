package solvekit

import (
	"fmt"
	"strings"

	parsekit "toolKit/parseKit"
)

type Ant struct {
	ID       int
	Path     []*parsekit.Room
	Step     int
	Finished bool
}

func Simulate(antCount int, start, end *parsekit.Room) {
	paths := FindDisjointPaths(start, end)
	if len(paths) == 0 {
		return
	}

	ants := make([]Ant, antCount)
	for i := 0; i < antCount; i++ {
		ants[i] = Ant{
			ID:   i + 1,
			Path: paths[i%len(paths)],
			Step: 0,
		}
	}

	turn := 0
	for {
		turn++
		var moves []string
		occupied := make(map[string]bool)



		for i := range ants {
			ant := &ants[i]
			if ant.Finished {
				continue
			}

			if ant.Step == len(ant.Path)-1 {
				ant.Finished = true
				continue
			}

			nextRoom := ant.Path[ant.Step+1]
			nextName := nextRoom.Name

			if nextName == end.Name || !occupied[nextName] {
				occupied[nextName] = true
				ant.Step++
				moves = append(moves, fmt.Sprintf("L%d-%s", ant.ID, nextName))
			}
		}

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

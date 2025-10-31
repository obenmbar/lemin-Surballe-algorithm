package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func MooveAnts(paths []Path, antNumber int, data string, assigned []int) {
	fmt.Println(strings.TrimSpace(data))
	fmt.Println("")

	ants := []Ant{}
	finished := 0
	ID := 1

	for finished < antNumber {
		mooves := []string{}

		for i := range ants {
			ant := &ants[i]
			last := len(ant.Path) - 1

			
			if ant.Position < last {
				ant.Position++
				mooves = append(mooves, ant.Name+"-"+ant.Path[ant.Position])

			}

			if ant.Position == last {
				finished++
			}
		}

		for i, path := range paths {
			if assigned[i] > 0 {
				newAnt := Ant{
					Name:     "L" + strconv.Itoa(ID),
					Path:     path,
					Position: 0,
				}

				ants = append(ants, newAnt)
				mooves = append(mooves, newAnt.Name+"-"+newAnt.Path[0])
				ID++
				assigned[i]--
			}
		}

		if len(mooves) > 0 {
			fmt.Println(strings.Join(mooves, " "))
		}
	}
}

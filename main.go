package main

import (
	"fmt"

	parsekit "toolKit/parseKit"
	solvekit "toolKit/solveKit"
)

func main() {
	lines, err := parsekit.ReadFileLines()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rest, err := parsekit.ParseAnts(lines)
	if err != nil {
		fmt.Println("Ant parsing error:", err)
		return
	}

	if err := parsekit.ParseRooms(rest); err != nil {
		fmt.Println("Room parsing error:", err)
		return
	}

	if err := parsekit.ParseStartEnd(rest); err != nil {
		fmt.Println("Start/End parsing error:", err)
		return
	}

	if err := parsekit.ParseTunnels(rest); err != nil {
		fmt.Println("Tunnel parsing error:", err)
		return
	}

	// if err := parsekit.WriteParsedOutput("parsed_output.txt"); err != nil {
	// 	fmt.Println("Error writing output file:", err)
	// 	return
	// }

	start := parsekit.Rooms[parsekit.StartRoom]
	end := parsekit.Rooms[parsekit.EndRoom]

	paths := solvekit.FindDisjointPaths(start, end)

	// if len(paths) == 0 {
	// 	fmt.Println("No path!")
	// 	return
	// }

	for i, path := range paths {
		fmt.Printf("Path %d: ", i+1)
		for j, room := range path {
			if j > 0 {
				fmt.Print(" -> ")
			}
			fmt.Print(room.Name)
		}
		fmt.Println()
	}
	solvekit.Simulate(parsekit.AntNum, start, end)
}

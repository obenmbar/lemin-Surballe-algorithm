package parsekit

import (
	"fmt"
	"os"
	"strings"
)

func IsEmpty(line string) bool {
	return strings.TrimSpace(line) == ""
}

func IsComment(line string) bool {
	line = strings.TrimSpace(line)
	return strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##")
}

func WriteParsedOutput(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "Ants: %d\n", AntNum)
	fmt.Fprintf(file, "Start Room: %s\n", StartRoom)
	fmt.Fprintf(file, "End Room: %s\n", EndRoom)

	fmt.Fprintln(file, "\nRooms:")
	for name, room := range Rooms {
		fmt.Fprintf(file, "  %s (%d, %d)\n", name, room.X, room.Y)
	}

	// fmt.Fprintln(file, "\nTunnels:")
	// for _, tunnel := range Tunnels {
	// 	fmt.Fprintf(file, "%s-%s\n", tunnel.From, tunnel.To)
	// }

	fmt.Fprintln(file, "\n Parsing completed successfully.")

	return nil
}

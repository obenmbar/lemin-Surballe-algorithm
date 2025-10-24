package parsekit

import (
	"fmt"
	"strings"
)

func ParseTunnels(lines []string) error {
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if !strings.Contains(line, "-") {
			continue
		}

		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return fmt.Errorf("invalid tunnel line: %s", line)
		}

		from := parts[0]
		to := parts[1]

		roomFrom, okFrom := Rooms[from]
		roomTo, okTo := Rooms[to]

		if !okFrom || !okTo {
			return fmt.Errorf("invalid tunnel: %s (one of the rooms doesn't exist)", line)
		}

		if from == to {
			return fmt.Errorf("invalid tunnel: %s cannot link to itself", from)
		}

		roomFrom.Link = append(roomFrom.Link, roomTo)
		roomTo.Link = append(roomTo.Link, roomFrom)

	}

	return nil
}

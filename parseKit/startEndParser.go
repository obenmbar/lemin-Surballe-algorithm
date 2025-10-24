package parsekit

import (
	"fmt"
	"strings"
)

func ParseStartEnd(lines []string) error {
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "##start" {
			for j := i + 1; j < len(lines); j++ {
				next := strings.TrimSpace(lines[j])
				if next == "" || strings.HasPrefix(next, "#") {
					continue
				}
				name, _, _, err := parseRoomDataStrict(next)
				if err != nil {
					return fmt.Errorf("invalid room after ##start at line %d: %v", j+1, err)
				}
				StartRoom = name
				break
			}
		} else if line == "##end" {
			for j := i + 1; j < len(lines); j++ {
				next := strings.TrimSpace(lines[j])
				if next == "" || strings.HasPrefix(next, "#") {
					continue
				}
				name, _, _, err := parseRoomDataStrict(next)
				if err != nil {
					return fmt.Errorf("invalid room after ##end at line %d: %v", j+1, err)
				}
				EndRoom = name
				break
			}
		}
	}

	if StartRoom == "" {
		return fmt.Errorf("missing start room (no valid room after ##start)")
	}
	if EndRoom == "" {
		return fmt.Errorf("missing end room (no valid room after ##end)")
	}

	return nil
}

func parseRoomDataStrict(line string) (string, int, int, error) {
	var name string
	var x, y int
	n, err := fmt.Sscanf(line, "%s %d %d", &name, &x, &y)
	if err != nil || n != 3 {
		return "", 0, 0, fmt.Errorf("line is not a valid room: %q", line)
	}
	if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
		return "", 0, 0, fmt.Errorf("invalid room name: %s", name)
	}
	return name, x, y, nil
}

package parsekit

import (
	"fmt"
	"strings"
)

func ParseRooms(lines []string) error {
	tunnelsStarted := false

	for i, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		isTunnel := strings.Contains(line, "-") && !strings.Contains(line, " ")

		if isTunnel {
			tunnelsStarted = true
			continue
		}

		if tunnelsStarted {
			return fmt.Errorf("room definition after tunnel at line %d: %q", i+1, line)
		}

		var name string
		var x, y int

		n, err := fmt.Sscanf(line, "%s %d %d", &name, &x, &y)
		if err != nil || n != 3 {
			return fmt.Errorf("invalid room line at %d: %q", i+1, line)
		}

		if strings.HasPrefix(name, "L") {
			return fmt.Errorf("invalid room name (starts with 'L'): %s", name)
		}

		for _, r := range Rooms {
			if r.X == x && r.Y == y {
				return fmt.Errorf("duplicate coordinates for room %s: (%d,%d)", name, x, y)
			}
		}

		Rooms[name] = &Room{Name: name, X: x, Y: y}
	}

	return nil
}

package parsekit

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseAnts(lines []string) ([]string, error) {
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil || n <= 0 {
			return nil, fmt.Errorf("invalid ants number: %s", line)
		}

		AntNum = n
		return lines[i+1:], nil
	}

	return nil, fmt.Errorf("no ants number found")
}

package parsekit

import (
	"fmt"
	"os"
	"strings"
)

func ReadFileLines() ([]string, error) {
	if len(os.Args) != 2 {
		return nil, fmt.Errorf("$Usage: go run program.go testFile.txt")
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	return lines, nil
}

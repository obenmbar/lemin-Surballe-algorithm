	package functions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ValidateFormat checks the input file format, parses its data, and returns a ready-to-use Farm.
func ValidateFormat(data string) (Farm, error) {
	if strings.TrimSpace(data) == "" {
		return Farm{}, fmt.Errorf("the file is empty")
	}

	lines := strings.Split(data, "\n")
	farm := initializeFarm()

	startIdx, err := parseAntNumber(lines, &farm)
	if err != nil {
		return Farm{}, err
	}

	foundTunnels := false
	if err := parseRoomsAndTunnels(lines, startIdx, &farm, &foundTunnels); err != nil {
		return Farm{}, err
	}

	if err := validateSpecialRooms(&farm); err != nil {
		return Farm{}, err
	}

	return farm, nil
}

// initializeFarm creates an empty Farm struct with initialized maps.
func initializeFarm() Farm {
	return Farm{
		Rooms:        make(map[string]*Room),
		SpecialRooms: make(map[string]string),
		Tunnels:      make(map[string]bool),
		Edges:        make(map[string]Edge),
	}
}

// parseAntNumber reads the first valid line to set the total number of ants.
func parseAntNumber(lines []string, farm *Farm) (int, error) {
	startIndex := 0
	for i, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" || isComment(line) {
			continue
		}

		ants, err := strconv.Atoi(line)
		if err != nil || ants < 1 {
			return 0, fmt.Errorf("first line must be a positive number of ants")
		}

		farm.Antnumber = ants
		startIndex = i + 1
		break
		
	}
	return startIndex, nil
}

// parseRoomsAndTunnels processes all lines to register rooms and tunnels in the farm.
func parseRoomsAndTunnels(lines []string, startIdx int, farm *Farm, foundTunnels *bool) error {
	for i := startIdx; i < len(lines); i++ {

		line := strings.TrimSpace(lines[i])

		if isEmpty(line) {
			continue
		}

		if isComment(line) {
			switch line {
			case "##start", "##end":
				if err := handleSpecialComment(lines, &i, farm, foundTunnels, line); err != nil {
					return err
				}
			}
			continue
		}

		if msg, ok := isRoomOrTunnel(line, farm, foundTunnels); !ok {
			return errors.New(msg)
		}
	}

	return nil
}

// validateSpecialRooms ensures both start and end rooms are defined.
func validateSpecialRooms(farm *Farm) error {
	if _, found := farm.SpecialRooms["start"]; !found {
		return fmt.Errorf("no start room found")
	}

	if _, found := farm.SpecialRooms["end"]; !found {
		return fmt.Errorf("no end room found")
	}

	return nil
}

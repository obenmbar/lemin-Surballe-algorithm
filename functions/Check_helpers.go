package functions

import (
	"fmt"
	"strings"
)

// handleSpecialComment processes '##start' or '%23%23end' lines and links them to the right room.
func handleSpecialComment(lines []string, idx *int, farm *Farm, foundTunnels *bool, specialType string) error {
	var Index_start int
	if *idx+1 >= len(lines) {
		return fmt.Errorf("'%s' must not be the last line", specialType)
	}
	for v := *idx + 1; v < len(lines); v++ {
		value_after_end_start:= strings.TrimSpace(lines[v])
		if value_after_end_start == "" || strings.HasPrefix(value_after_end_start, "#") {
			if strings.TrimSpace(lines[v]) != "##start" &&  strings.TrimSpace(lines[v]) != "##end"{
				Index_start += 1
				continue
			} else {
				return fmt.Errorf("error il ya la valeur ##start directement desous de la valeur ##end ou le contraire ")
			}
		} else {
			break
		}
	}
	nextLine := strings.TrimSpace(lines[*idx+Index_start+1])

	if msg, ok := isRoomOrTunnel(nextLine, farm, foundTunnels); !ok {
		return fmt.Errorf("after '%s' the next line must be a valid room: %s", specialType, msg)
	}

	name := strings.Fields(nextLine)[0]

	switch specialType {

	case "##start":
		if _, exists := farm.SpecialRooms["start"]; exists {
			return fmt.Errorf("more than one '##start' found")
		}
		farm.SpecialRooms["start"] = name

	case "##end":
		if _, exists := farm.SpecialRooms["end"]; exists {
			return fmt.Errorf("more than one '##end' found")
		}
		farm.SpecialRooms["end"] = name
	}

	*idx += Index_start+1
	return nil
}

// isRoomOrTunnel checks if a given line defines a room, tunnel, or something invalid.
func isRoomOrTunnel(line string, farm *Farm, foundTunnel *bool) (string, bool) {
	if isEmpty(line) {
		return "line is empty", false
	}

	if isComment(line) {
		return "line is a comment", false
	}

	if isRoomLine(line) {
		return handleRoomLine(line, farm, foundTunnel)
	}

	if isTunnelLine(line) {
		return handleTunnelLine(line, farm, foundTunnel)
	}

	return fmt.Sprintf("if you want to comment, you need to use '#' at the begining of the line: %s", line), false
}

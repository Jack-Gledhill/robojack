package presence

import (
	"strconv"
	"strings"

	"google.golang.org/api/calendar/v3"
)

const (
	// CharacterDelimiter is the prefix for the character name in the event description
	CharacterDelimiter = "Character: "
	// SystemDelimiter is the prefix for the TTRPG system name in the event description
	SystemDelimiter = "System: "
	// PlayerCountDelimiter is the prefix for the player count in the event description
	PlayerCountDelimiter = "Players: "
)

// TTRPGMeta holds information about a TTRPG game that I'm currently playing
type TTRPGMeta struct {
	Character string
	Players   int
	System    string
}

// GetTTRPGMeta processes an event and returns any TTRPG metadata found in the description
func GetTTRPGMeta(e *calendar.Event) TTRPGMeta {
	meta := TTRPGMeta{
		Character: "Nameless Hero",
		System:    "DnD 5e 2014",
		Players:   3, // Me, a GM and a player
	}
	lines := strings.Split(e.Description, "\n")

	for _, l := range lines {
		if strings.HasPrefix(l, CharacterDelimiter) {
			meta.Character = strings.TrimPrefix(l, CharacterDelimiter)
		} else if strings.HasPrefix(l, SystemDelimiter) {
			meta.System = strings.TrimPrefix(l, SystemDelimiter)
		} else if strings.HasPrefix(l, PlayerCountDelimiter) {
			players, _ := strconv.Atoi(strings.TrimPrefix(l, PlayerCountDelimiter))
			meta.Players = players
		}
	}

	return meta
}

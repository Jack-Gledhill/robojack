package events

import "github.com/bwmarrin/discordgo"

// Handlers contains all the functions that will be registered as event handlers with discordgo
var Handlers []interface{}

// Register creates a new event handler and adds it to the Handlers slice
func Register(s *discordgo.Session) {
	for _, h := range Handlers {
		s.AddHandler(h)
	}
}

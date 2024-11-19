package events

import (
	"github.com/Jack-Gledhill/robojack/bot/commands"

	"github.com/bwmarrin/discordgo"
)

func init() {
	Handlers = append(Handlers, InteractionCreate)
}

// InteractionCreate listens for commands and calls the appropriate handler function
func InteractionCreate(s *discordgo.Session, e *discordgo.InteractionCreate) {
	if h, ok := commands.Handlers[e.ApplicationCommandData().Name]; ok {
		h(s, e)
	}
}

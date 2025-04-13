package commands

import "github.com/bwmarrin/discordgo"

// Handlers contains a mapping of command names to their respective handler functions
var Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}

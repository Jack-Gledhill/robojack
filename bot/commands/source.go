package commands

import (
	"fmt"

	"github.com/Jack-Gledhill/robojack/debug"

	"github.com/bwmarrin/discordgo"
)

func init() {
	Handlers["source"] = Source
}

// Source just returns a link to the bot's GitHub repository, nothing fancy
func Source(s *discordgo.Session, e *discordgo.InteractionCreate) {
	s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "🧑🏻‍💻 Source Code",
					Description: fmt.Sprintf("You know it's rude to ask a robot about his source code right? Well if you must, my repo is on [GitHub](%s).", debug.Git.Repository),
				},
			},
		},
	})
}

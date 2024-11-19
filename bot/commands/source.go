package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// SourceRepository is the base URL for the bot's repository
const SourceRepository = "https://github.com/Jack-Gledhill/robojack"

func init() {
	New("source", "Get the link to my source code", Source)
}

// Source just returns a link to the bot's GitHub repository, nothing fancy
func Source(s *discordgo.Session, e *discordgo.InteractionCreate) {
	s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "🧑🏻‍💻 Source Code",
					Description: fmt.Sprintf("You know it's rude to ask a robot about his source code right? Well if you must, my repo is on [GitHub](%s).", SourceRepository),
				},
			},
		},
	})
}

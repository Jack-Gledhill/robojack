package commands

import (
	"fmt"

	"github.com/Jack-Gledhill/robojack/env"

	"github.com/bwmarrin/discordgo"
)

func init() {
	New("token", "Gets the bot's token", Token)
}

// Token DMs the bot's token to the owner, but mocks anyone else who attempts to use the command
func Token(s *discordgo.Session, e *discordgo.InteractionCreate) {
	// They're not the owner, so let's mock them for fun
	if e.Member.User.ID != env.Bot.OwnerID {
		err := s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("||%s HA! Wow, you really thought I was gonna tell *you*? Fuck no.||", env.Bot.Emojis.MiddleFinger),
			},
		})
		if err != nil {
			panic(err)
		}
		return
	}

	err := s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: fmt.Sprintf("%s I gotcha bud, here's my token: ||%s||", env.Bot.Emojis.Wink, env.Bot.Token),
		},
	})
	if err != nil {
		panic(err)
	}
}

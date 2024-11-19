package commands

import (
	"fmt"

	"github.com/Jack-Gledhill/robojack/env"

	"github.com/bwmarrin/discordgo"
)

// InviteLinkBase is the base URL for the invite link
const InviteLinkBase = "https://discord.com/oauth2/authorize?client_id=%s&scope=applications.commands%%20bot"

func init() {
	New("invite", "Get a link to invite me to your server", Invite)
}

// Invite returns a link to invite the bot to a server, assuming the user is the owner of the bot
func Invite(s *discordgo.Session, e *discordgo.InteractionCreate) {
	// User is the owner so send the invite link
	if e.Member.User.ID == env.Bot.OwnerID {
		url := fmt.Sprintf(InviteLinkBase, s.State.User.ID)
		err := s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("%s I gotcha bud, invite me [here](%s)", env.Bot.Emojis.Wink, url),
			},
		})
		if err != nil {
			panic(err)
		}

		return
	}

	// User isn't owner, mock them instead
	err := s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("%s Sorry hon, but you just ain't good enough for me. Better luck next time.", env.Bot.Emojis.MiddleFinger),
		},
	})
	if err != nil {
		panic(err)
	}
}

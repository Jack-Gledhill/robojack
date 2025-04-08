package commands

import (
	"math/rand"

	"github.com/Jack-Gledhill/robojack/bot/insult"

	"github.com/bwmarrin/discordgo"
)

func init() {
	New("insult", "Will randomly insult you", Insult)
}

func Insult(s *discordgo.Session, e *discordgo.InteractionCreate) {
	// TODO: attach this to a database with flags per server
	insults := insult.DefaultList()

	err := s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: insults[rand.Intn(len(insults))],
		},
	})
	if err != nil {
		panic(err)
	}
}

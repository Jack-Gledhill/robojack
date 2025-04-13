package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	Handlers["ping"] = Ping
}

// Ping calculates and presents the bot's latencies to the Discord HTTP & WebSocket APIs.
func Ping(s *discordgo.Session, e *discordgo.InteractionCreate) {
	// TODO: pull down Discord status?

	// Send the initial message containing the heartbeat latency
	before := time.Now()
	s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "🏓 Pong!",
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "☁️ WebSocket Latency",
							Value: fmt.Sprintf("`%d ms`", s.HeartbeatLatency().Milliseconds()),
						},
						{
							Name:  "💬 HTTP Latency",
							Value: "⏳ Pinging...",
						},
					},
				},
			},
		},
	})
	after := time.Now()

	// Edit the message to include the REST API latency
	s.InteractionResponseEdit(e.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{
			{
				Title: "🏓 Pong!",
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "☁️ WebSocket Latency",
						Value: fmt.Sprintf("`%d ms`", s.HeartbeatLatency().Milliseconds()),
					},
					{
						Name:  "💬 HTTP Latency",
						Value: fmt.Sprintf("`%d ms`", after.Sub(before).Milliseconds()),
					},
				},
			},
		},
	})
}

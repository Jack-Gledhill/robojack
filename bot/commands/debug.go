package commands

import (
	"fmt"
	"math"
	"runtime"

	"github.com/Jack-Gledhill/robojack/env"
	"github.com/gin-gonic/gin"

	"github.com/bwmarrin/discordgo"
)

func init() {
	New("debug", "Shows some information about the running environment", Debug)
}

// Debug returns some basic information about the environment the bot is running in
func Debug(s *discordgo.Session, e *discordgo.InteractionCreate) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	err := s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "🔧 Debugging Information",
					Description: fmt.Sprintf("Environment: `%s`\nOperating system: `%s`\nRegistered owner: <@%s>", env.Mode(), env.Build.OS, env.Bot.OwnerID),
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "Versions",
							Value: fmt.Sprintf("Go: `%s`\nDiscordGo: `%s`\nGin: `%s`\nCommit Hash: %s", runtime.Version(), discordgo.VERSION, gin.Version, GetCommit()),
						},
						{
							Name:  "Memory",
							Value: fmt.Sprintf("Allocated: `%d MB`\nReserved: `%d MB`\nGC cycles: `%d`", BtoMB(m.Alloc), BtoMB(m.Sys), m.NumGC),
						},
						{
							Name:  "CPU",
							Value: fmt.Sprintf("Logical Cores: `%d`\nActive Goroutines: `%d`\nGarbage Collector usage: `%d%%`", runtime.NumCPU(), runtime.NumGoroutine(), ToPercent(m.GCCPUFraction)),
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}

func GetCommit() string {
	if env.Build.Git.Modified {
		return fmt.Sprintf("[%s](%s/commit/%s) (modified)", env.Build.Git.Revision[:7], SourceRepository, env.Build.Git.Revision)
	}

	return fmt.Sprintf("[%s](%s)", env.Build.Git.Revision[:7], SourceRepository)
}

func BtoMB(b uint64) uint64 {
	return b / 1000 / 1000
}

func ToPercent(f float64) int {
	return int(math.Round(f * 100))
}

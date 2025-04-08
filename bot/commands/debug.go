package commands

import (
	"fmt"
	"math"
	"runtime"

	"github.com/Jack-Gledhill/robojack/config"

	"github.com/bwmarrin/discordgo"
)

const (
	basicInfo   = "Environment: `%s`\nOperating System: `%s`\nRegistered Owner: <@%s>"
	cpuInfo     = "Logical Cores: `%d`\nActive Goroutines: `%d`\nGC Usage: `%d%%`"
	gitInfo     = "Commit: %s\nBranch/tag: `%s`"
	memoryInfo  = "Allocated: `%d MB`\nReserved: `%d MB`\nGC Cycles: `%d`"
	versionInfo = "Go: `%s`\nDiscordGo: `%s`"
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
					Description: fmt.Sprintf(basicInfo, config.Mode(), config.Build.OS, config.Bot.OwnerID),
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "Versions",
							Value: fmt.Sprintf(versionInfo, runtime.Version(), discordgo.VERSION),
						},
						{
							Name:  "Git",
							Value: fmt.Sprintf(gitInfo, GetCommit(), config.Git.Ref),
						},
						{
							Name:  "Memory",
							Value: fmt.Sprintf(memoryInfo, BtoMB(m.Alloc), BtoMB(m.Sys), m.NumGC),
						},
						{
							Name:  "CPU",
							Value: fmt.Sprintf(cpuInfo, runtime.NumCPU(), runtime.NumGoroutine(), ToPercent(m.GCCPUFraction)),
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
	if config.Git.Revision == "" {
		return "unknown"
	}

	if config.Git.Modified {
		return fmt.Sprintf("[%s](%s/commit/%s) (modified)", config.Git.Revision[:7], config.Git.Repository, config.Git.Revision)
	}

	return fmt.Sprintf("[%s](%s)", config.Git.Revision[:7], config.Git.Repository)
}

func BtoMB(b uint64) uint64 {
	return b / 1000 / 1000
}

func ToPercent(f float64) int {
	return int(math.Round(f * 100))
}

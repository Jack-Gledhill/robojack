package debug

import (
	"runtime"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

// Build is populated at runtime with version information
var Build = BuildInfo{
	DiscordGo: discordgo.VERSION,
	Gin:       gin.Version,
	Go:        runtime.Version(),
}

// BuildInfo holds information about the packages this software was built with
type BuildInfo struct {
	DiscordGo string `json:"discordgo"`
	Gin       string `json:"gin"`
	Go        string `json:"golang"`
}

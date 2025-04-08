package config

import "github.com/bwmarrin/discordgo"

var Bot *BotConfig

func init() {
	Bot = &file.Bot
}

type BotConfig struct {
	Mode    string `yaml:"mode"`
	OwnerID string `yaml:"owner_id"`
	Token   string `yaml:"token"`
}

// Owner fetches the bot owner from Discord and returns information about them
func (b *BotConfig) Owner(s *discordgo.Session) (*discordgo.User, error) {
	return s.User(b.OwnerID)
}

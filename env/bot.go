package env

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

const (
	// BotPrefix is the prefix used for bot-related environment variables
	BotPrefix = "DISCORD_"
	// EmojiPrefix is the prefix used for bot emoji environment variables
	EmojiPrefix = BotPrefix + "EMOJI_"
)

// Bot is an instance of BotConfig, holding most of the config options available
var Bot = BotConfig{
	Emojis: BotEmojis{
		Angry:        os.Getenv(EmojiPrefix + "ANGRY"),
		Crying:       os.Getenv(EmojiPrefix + "CRYING"),
		Dead:         os.Getenv(EmojiPrefix + "DEAD"),
		Facepalm:     os.Getenv(EmojiPrefix + "FACEPALM"),
		MiddleFinger: os.Getenv(EmojiPrefix + "MIDDLE_FINGER"),
		LaserEyes:    os.Getenv(EmojiPrefix + "LASER_EYES"),
		Love:         os.Getenv(EmojiPrefix + "LOVE"),
		Peace:        os.Getenv(EmojiPrefix + "PEACE"),
		Peeking:      os.Getenv(EmojiPrefix + "PEEKING"),
		Smiling:      os.Getenv(EmojiPrefix + "SMILING"),
		Sunglasses:   os.Getenv(EmojiPrefix + "SUNGLASSES"),
		Thinking:     os.Getenv(EmojiPrefix + "THINKING"),
		Waving:       os.Getenv(EmojiPrefix + "WAVING"),
		Weary:        os.Getenv(EmojiPrefix + "WEARY"),
		Wink:         os.Getenv(EmojiPrefix + "WINK"),
	},
	OwnerID: os.Getenv(BotPrefix + "OWNER_ID"),
	Token:   RequiredVar(BotPrefix + "TOKEN"),
}

// BotConfig holds all config options that apply to Discord and the bot in general
type BotConfig struct {
	Emojis  BotEmojis
	OwnerID string
	Token   string
}

// Owner fetches the bot owner from Discord and returns information about them
func (b *BotConfig) Owner(s *discordgo.Session) (*discordgo.User, error) {
	return s.User(b.OwnerID)
}

// BotEmojis are used as decoration in bot responses
type BotEmojis struct {
	Angry        string
	Crying       string
	Dead         string
	Facepalm     string
	MiddleFinger string
	LaserEyes    string
	Love         string
	Peace        string
	Peeking      string
	Smiling      string
	Sunglasses   string
	Thinking     string
	Waving       string
	Weary        string
	Wink         string
}

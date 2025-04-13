package bot

import (
	"github.com/Jack-Gledhill/robojack/bot/events"
	"github.com/Jack-Gledhill/robojack/config"
	"github.com/Jack-Gledhill/robojack/log"

	"github.com/bwmarrin/discordgo"
)

// Session is the discordgo.Session used to connect to Discord
var Session *discordgo.Session

var (
	kill  = make(chan bool, 1)
	ready = make(chan bool, 1)
)

func init() {
	// Initialise the bot
	var err error
	Session, err = discordgo.New("Bot " + config.Bot.Token)
	if err != nil {
		panic(err)
	}

	// Set up the bot's identify properties
	Session.Identify.Intents = discordgo.IntentsGuildMessages
	Session.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{
			Name:  "custom",
			State: "YIPPEE",
			Type:  discordgo.ActivityTypeCustom,
		},
		Status: "dnd",
		AFK:    false,
	}

	// Register event handlers
	Session.AddHandlerOnce(func(s *discordgo.Session, e *discordgo.Ready) {
		log.Info("Bot is now READY")
		ready <- true
	})

	events.Register(Session)
}

// Start will connect to Discord and block until KillSwitch receives a boolean
func Start() {
	log.Debug("Connecting to Discord...")

	// Start the bot
	err := Session.Open()
	if err != nil {
		panic(err)
	}

	// Block until a kill signal is received
	<-kill
}

// Close will stop the bot and deregister any commands it had
func Close() {
	log.Debug("Bot is stopping...")
	err := Session.Close()
	if err != nil {
		panic(err)
	}

	// Terminate the Start() goroutine
	kill <- true
	log.Warn("Bot has stopped")
}

// WaitUntilReady blocks the current thread until the bot has received the READY event from Discord
// If the bot is already ready when this function is called, it will return immediately
func WaitUntilReady() {
	<-ready
}

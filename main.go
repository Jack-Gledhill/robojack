package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Jack-Gledhill/robojack/bot/commands"
	"github.com/Jack-Gledhill/robojack/bot/events"
	"github.com/Jack-Gledhill/robojack/config"

	"github.com/bwmarrin/discordgo"
)

var bot *discordgo.Session

func init() {
	// Initialise the bot
	var err error
	bot, err = discordgo.New("Bot " + config.Bot.Token)
	if err != nil {
		panic(err)
	}

	// Set up the bot's identify properties
	bot.Identify.Intents = discordgo.IntentsGuildMessages
	bot.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{
			Name:  "custom",
			State: "YIPPEE",
			Type:  discordgo.ActivityTypeCustom,
		},
		Status: "dnd",
		AFK:    false,
	}
}

func main() {
	// Recover from any panics and log them
	// This handler isn't added until after init() completes, so any panics in init() will still crash the program
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from a panic: ", err)
		}
	}()

	// Start the bot
	// This is NOT a blocking call, as the websocket is run in a new goroutine
	err := bot.Open()
	if err != nil {
		panic(err)
	}
	defer bot.Close()

	// Register any commands and event handlers
	events.Register(bot)
	commands.Register(bot)
	defer commands.Deregister(bot)

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}

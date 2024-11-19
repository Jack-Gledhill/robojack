package main

import (
	"fmt"

	"github.com/Jack-Gledhill/robojack/bot/commands"
	"github.com/Jack-Gledhill/robojack/bot/events"
	"github.com/Jack-Gledhill/robojack/env"
	"github.com/Jack-Gledhill/robojack/web"

	"github.com/bwmarrin/discordgo"
)

// Bot is the discordgo.Session that connects the program to Discord
var Bot *discordgo.Session

func init() {
	// Initialise the bot
	var err error
	Bot, err = discordgo.New("Bot " + env.Bot.Token)
	if err != nil {
		panic(err)
	}

	// Set up the bot's identify properties
	Bot.Identify.Intents = discordgo.IntentsGuildMessages
	Bot.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{
			Name:  "custom",
			State: "YIIPPEE",
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
	// This is NOT a blocking call, as the websocket is ran in a new goroutine
	err := Bot.Open()
	if err != nil {
		panic(err)
	}
	defer Bot.Close()

	// Register any commands and event handlers
	events.Register(Bot)
	commands.Register(Bot)
	defer commands.Deregister(Bot)

	// Start the web server
	// This is a blocking call, and will keep the bot running too
	web.Server.Run(env.Web.ListenAddress())
}

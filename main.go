package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/Jack-Gledhill/robojack/bot/commands"
	"github.com/Jack-Gledhill/robojack/bot/events"
	"github.com/Jack-Gledhill/robojack/config"
	"github.com/Jack-Gledhill/robojack/log"

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
			log.L.Error().Msg(fmt.Sprintf("Recovered from a panic: %s", err))
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

	// Print some info to indicate that the bot has started up
	log.Info(" _____   ____  ____   ____       _         _____ _  __")
	log.Info("|  __ \\ / __ \\|  _ \\ / __ \\     | |  /\\   / ____| |/ /")
	log.Info("| |__) | |  | | |_) | |  | |    | | /  \\ | |    | ' / ")
	log.Info("|  _  /| |  | |  _ <| |  | |_   | |/ /\\ \\| |    |  <  ")
	log.Info("| | \\ \\| |__| | |_) | |__| | |__| / ____ \\ |____| . \\ ")
	log.Info("|_|  \\_\\\\____/|____/ \\____/ \\____/_/    \\_\\_____|_|\\_\\")
	log.Info("===== Build Information =====")
	log.Info("DiscordGo: %s", discordgo.VERSION)
	log.Info("Go:        %s", runtime.Version())
	log.Info("OS:        %s", config.Build.OS)
	log.Info("Revision:  %s", config.Git.Revision)
	log.Info("Version:   %s", config.Git.Ref)
	log.Info("=============================")

	// Block until a kill signal is received
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Info("===== Shutting down =====")
}

package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/Jack-Gledhill/robojack/bot"
	"github.com/Jack-Gledhill/robojack/config"
	"github.com/Jack-Gledhill/robojack/log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	// Recover from any panics and log them
	defer func() {
		if err := recover(); err != nil {
			log.Error("Recovered from a panic: %s", err)
		}
	}()
}

func main() {
	go bot.Start()
	defer bot.Close()

	bot.WaitUntilReady()

	// Print some info to indicate that the bot has started up
	log.Info(" _____   ____  ____   ____       _         _____ _  __")
	log.Info("|  __ \\ / __ \\|  _ \\ / __ \\     | |  /\\   / ____| |/ /")
	log.Info("| |__) | |  | | |_) | |  | |    | | /  \\ | |    | ' / ")
	log.Info("|  _  /| |  | |  _ <| |  | |_   | |/ /\\ \\| |    |  <  ")
	log.Info("| | \\ \\| |__| | |_) | |__| | |__| / ____ \\ |____| . \\ ")
	log.Info("|_|  \\_\\\\____/|____/ \\____/ \\____/_/    \\_\\_____|_|\\_\\")
	log.Info("===== Build & Runtime Information =====")
	log.Info("Mode:      %s", config.Mode())
	log.Info("Log Level: %s", log.Level.String())
	log.Info("DiscordGo: %s", discordgo.VERSION)
	log.Info("Go:        %s", runtime.Version())
	log.Info("OS:        %s", config.Build.OS)
	log.Info("Revision:  %s", config.Git.Revision)
	log.Info("Version:   %s", config.Git.Ref)
	log.Info("=======================================")

	// Block until a kill signal is received
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Info("===== Shutting down =====")
}

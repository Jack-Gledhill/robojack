package main

import (
	"github.com/Jack-Gledhill/robojack/bot"
	"github.com/Jack-Gledhill/robojack/debug"
	"github.com/Jack-Gledhill/robojack/log"
	"github.com/Jack-Gledhill/robojack/web"
)

func init() {
	// Recover from any panics and log them
	defer func() {
		if err := recover(); err != nil {
			log.Error("Recovered from a panic: %s", err)
		}
	}()

	log.Info(" _____   ____  ____   ____       _         _____ _  __")
	log.Info("|  __ \\ / __ \\|  _ \\ / __ \\     | |  /\\   / ____| |/ /")
	log.Info("| |__) | |  | | |_) | |  | |    | | /  \\ | |    | ' / ")
	log.Info("|  _  /| |  | |  _ <| |  | |_   | |/ /\\ \\| |    |  <  ")
	log.Info("| | \\ \\| |__| | |_) | |__| | |__| / ____ \\ |____| . \\ ")
	log.Info("|_|  \\_\\\\____/|____/ \\____/ \\____/_/    \\_\\_____|_|\\_\\")
	log.Info("===== Build & Runtime Information =====")
	log.Info("Mode:      %s", debug.Runtime.Mode)
	log.Info("Log Level: %s", debug.Runtime.LogLevel)
	log.Info("DiscordGo: %s", debug.Build.DiscordGo)
	log.Info("Gin:       %s", debug.Build.Gin)
	log.Info("Go:        %s", debug.Build.Go)
	log.Info("OS:        %s", debug.System.OS)
	log.Info("Arch:      %s", debug.System.Arch)
	log.Info("Revision:  %s", debug.Git.Commit.Hash)
	log.Info("Version:   %s", debug.Git.Ref)
	log.Info("=======================================")
}

func main() {
	defer log.Info("===== Shutting down =====")

	// Start the bot in a separate goroutine
	go bot.Start()
	defer bot.Close()

	// Start the webserver, this will block until the program exits
	web.Start()
}

package commands

import (
	"github.com/Jack-Gledhill/robojack/log"

	"github.com/bwmarrin/discordgo"
)

var (
	// Commands contains all the commands that will be registered when the bot starts
	Commands []*discordgo.ApplicationCommand
	// Handlers contains a mapping of command names to their respective handler functions
	Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
	// Registered contains all the commands that have been registered with Discord when the bot starts
	Registered []*discordgo.ApplicationCommand
)

// Register sends the list of Commands to Discord to be registered
func Register(s *discordgo.Session) {
	for _, c := range Commands {
		log.Debug("Registering command: %s", c.Name)

		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", c)
		if err != nil {
			panic(err)
		}

		Registered = append(Registered, cmd)
	}
}

// Deregister is the inverse of Register and should be called when the bot exits
func Deregister(s *discordgo.Session) {
	for _, c := range Registered {
		log.Debug("Deregistering command: %s", c.Name)

		err := s.ApplicationCommandDelete(s.State.User.ID, "", c.ID)
		if err != nil {
			panic(err)
		}
	}

	Registered = nil
}

// New sets up a new command and adds it to Handlers
func New(name string, description string, callback func(s *discordgo.Session, e *discordgo.InteractionCreate)) {
	Commands = append(Commands, &discordgo.ApplicationCommand{
		Name:        name,
		Description: description,
	})
	Handlers[name] = callback
}

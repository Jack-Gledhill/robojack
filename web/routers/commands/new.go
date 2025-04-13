package commands

import (
	"net/http"

	"github.com/Jack-Gledhill/robojack/bot"
	"github.com/Jack-Gledhill/robojack/log"
	"github.com/Jack-Gledhill/robojack/web/response"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func add(c *gin.Context) {
	var data discordgo.ApplicationCommand
	err := c.BindJSON(&data)
	if err != nil {
		log.Error("Error reading request body: %s", err.Error())

		response.New().
			Status(http.StatusBadRequest).
			Msg("malformed request body").
			Send(c)
		return
	}

	// Attempt to create the command
	_, err = bot.Session.ApplicationCommandCreate(bot.Session.State.User.ID, "", &data)
	if err != nil {
		log.Error("Error creating command: %s", err.Error())

		response.New().
			Status(http.StatusInternalServerError).
			Msg("error when creating command").
			Send(c)
		return
	}

	response.New().
		Status(http.StatusCreated).
		Send(c)
}

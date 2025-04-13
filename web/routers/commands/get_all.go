package commands

import (
	"net/http"

	"github.com/Jack-Gledhill/robojack/bot"
	"github.com/Jack-Gledhill/robojack/log"
	"github.com/Jack-Gledhill/robojack/web/response"

	"github.com/gin-gonic/gin"
)

func getAll(c *gin.Context) {
	cmds, err := bot.Session.ApplicationCommands(bot.Session.State.User.ID, "")
	if err != nil {
		log.Error("Error fetching application commands: %s", err)

		response.New().
			Status(http.StatusInternalServerError).
			Msg("could not fetch commands").
			Send(c)
		return
	}

	response.New().
		Data(cmds).
		Send(c)
}

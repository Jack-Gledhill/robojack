package commands

import (
	"net/http"

	"github.com/Jack-Gledhill/robojack/bot"
	"github.com/Jack-Gledhill/robojack/log"
	"github.com/Jack-Gledhill/robojack/web/response"

	"github.com/gin-gonic/gin"
)

func getOne(c *gin.Context) {
	id := c.Param("id")

	cmd, err := bot.Session.ApplicationCommand(bot.Session.State.User.ID, "", id)
	if err != nil {
		log.Error("Error getting command: %s", err.Error())

		response.New().
			Status(http.StatusInternalServerError).
			Msg("error fetching command").
			Send(c)
		return
	}

	response.New().
		Data(cmd).
		Send(c)
}

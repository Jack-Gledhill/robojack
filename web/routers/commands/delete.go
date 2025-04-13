package commands

import (
	"net/http"

	"github.com/Jack-Gledhill/robojack/bot"
	"github.com/Jack-Gledhill/robojack/log"
	"github.com/Jack-Gledhill/robojack/web/response"

	"github.com/gin-gonic/gin"
)

func remove(c *gin.Context) {
	id := c.Param("id")

	err := bot.Session.ApplicationCommandDelete(bot.Session.State.User.ID, "", id)
	if err != nil {
		log.Error("Error deleting command: %s", err.Error())

		response.New().
			Status(http.StatusInternalServerError).
			Msg("error when deleting command").
			Send(c)
	}

	c.Status(http.StatusNoContent)
}

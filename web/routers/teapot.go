package routers

import (
	"net/http"

	"github.com/Jack-Gledhill/robojack/web/response"

	"github.com/gin-gonic/gin"
)

func teapot(c *gin.Context) {
	response.New().
		Status(http.StatusTeapot).
		Msg("I'm a little teapot,\nShort and stout,\nHere is my handle\nHere is my spout\nWhen I get all steamed up,\nHear me shout,\nTip me over and pour me out!").
		Send(c)
}

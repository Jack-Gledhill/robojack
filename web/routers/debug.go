package routers

import (
	"github.com/Jack-Gledhill/robojack/debug"
	"github.com/Jack-Gledhill/robojack/web/response"

	"github.com/gin-gonic/gin"
)

func Debug(c *gin.Context) {
	response.New().
		Data(gin.H{
			"build":   debug.Build,
			"git":     debug.Git,
			"runtime": debug.Runtime.Snapshot(),
			"system":  debug.System,
		}).
		Send(c)
}

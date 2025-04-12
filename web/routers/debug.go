package routers

import (
	dbg "github.com/Jack-Gledhill/robojack/debug"
	"github.com/Jack-Gledhill/robojack/web/response"

	"github.com/gin-gonic/gin"
)

func debug(c *gin.Context) {
	response.New().
		Data(gin.H{
			"build":   dbg.Build,
			"git":     dbg.Git,
			"runtime": dbg.Runtime.Snapshot(),
			"system":  dbg.System,
		}).
		Send(c)
}

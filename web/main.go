package web

import (
	"github.com/Jack-Gledhill/robojack/env"

	"github.com/gin-gonic/gin"
)

// Server is the web server used for the bot's API
var Server *gin.Engine

func init() {
	// Disable Gin's debug mode in production
	if env.Production() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create a new Gin server
	Server = gin.New()
}

package web

import (
	"github.com/Jack-Gledhill/robojack/config"
	"github.com/Jack-Gledhill/robojack/log"
	"github.com/Jack-Gledhill/robojack/web/routers"
	"github.com/Jack-Gledhill/robojack/web/routers/auth"

	"github.com/gin-gonic/gin"
)

var engine = gin.New()

func init() {
	// Disables Gin's verbose logging in production
	if config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Add Gin's default panic & 500 error recovery middleware
	engine.Use(gin.Recovery())

	// Add all routers to the engine
	routers.AddHandlers(engine.Group(""))
	auth.AddHandlers(engine.Group("/auth"))
}

// Start will call engine.Run() and block until gin encounters an error or the program terminates
func Start() {
	log.Info("Web server is listening on %s", config.Web.Port)
	err := engine.Run(config.Web.Port)
	if err != nil {
		panic(err)
	}
}

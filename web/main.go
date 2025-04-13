package web

import (
	"github.com/Jack-Gledhill/robojack/config"
	"github.com/Jack-Gledhill/robojack/log"
	"github.com/Jack-Gledhill/robojack/web/routers"
	"github.com/Jack-Gledhill/robojack/web/routers/auth"
	"github.com/Jack-Gledhill/robojack/web/routers/commands"
	"github.com/Jack-Gledhill/robojack/web/routers/dashboard"

	"github.com/a-h/templ/examples/integration-gin/gintemplrenderer"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	// Disables Gin's verbose logging in production
	if config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	engine = gin.New()

	// Silence trusted proxies warning
	_ = engine.SetTrustedProxies(nil)

	// Setup templ renderer
	defaultRenderer := engine.HTMLRender
	engine.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{
		FallbackHtmlRenderer: defaultRenderer,
	}

	// Add Gin's default panic & 500 error recovery middleware
	engine.Use(gin.Recovery())

	// Expose static files
	engine.Static("/static", "web/src/static")

	// Add all routers to the engine
	routers.AddHandlers(engine.Group(""))
	auth.AddHandlers(engine.Group("/auth"))
	commands.AddHandlers(engine.Group("/commands"))
	dashboard.AddHandlers(engine.Group("/dashboard"))
}

// Start will call engine.Run() and block until gin encounters an error or the program terminates
func Start() {
	log.Info("Web server is listening on %s", config.Web.Port)
	err := engine.Run(config.Web.Port)
	if err != nil {
		panic(err)
	}
}

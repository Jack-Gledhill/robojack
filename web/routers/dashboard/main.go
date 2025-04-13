package dashboard

import (
	"github.com/Jack-Gledhill/robojack/web/middleware"

	"github.com/gin-gonic/gin"
)

func AddHandlers(g *gin.RouterGroup) {
	g.Use(middleware.AuthenticationWithRedirect)
	g.GET("", index)
}

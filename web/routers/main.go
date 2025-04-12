package routers

import "github.com/gin-gonic/gin"

func AddHandlers(g *gin.RouterGroup) {
	g.GET("/debug", debug)
	g.GET("/ping", ping)
	g.GET("/teapot", teapot)
}

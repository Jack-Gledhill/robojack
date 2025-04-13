package commands

import (
	"github.com/Jack-Gledhill/robojack/web/middleware"

	"github.com/gin-gonic/gin"
)

func AddHandlers(g *gin.RouterGroup) {
	g.Use(middleware.Authentication)

	g.GET("/", getAll)
	g.GET("/:id", getOne)

	g.POST("/", add)
	g.DELETE("/:id", remove)
}

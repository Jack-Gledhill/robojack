package auth

import (
	"github.com/Jack-Gledhill/robojack/web/middleware"
	"github.com/gin-gonic/gin"
)

func AddHandlers(g *gin.RouterGroup) {
	g.GET("/login", login)
	g.GET("/callback", callback)
	g.GET("/me", middleware.Authentication, me)
	g.GET("/logout", logout)
}

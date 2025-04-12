package auth

import "github.com/gin-gonic/gin"

const (
	cookieToken = "authorization"
	queryCode   = "code"
	queryState  = "state"
)

func AddHandlers(g *gin.RouterGroup) {
	g.GET("/login", login)
	g.GET("/callback", callback)
	g.GET("/me", me)
	g.GET("/logout", logout)
}

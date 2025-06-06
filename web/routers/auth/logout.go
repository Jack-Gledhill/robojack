package auth

import (
	"net/http"

	"github.com/Jack-Gledhill/robojack/config"

	"github.com/gin-gonic/gin"
)

func logout(c *gin.Context) {
	c.SetCookie(config.Web.JWT.Cookie, "", -1, "/", config.Web.Domain().Hostname(), false, false)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

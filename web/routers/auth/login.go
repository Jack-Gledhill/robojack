package auth

import (
	"net/http"

	"github.com/Jack-Gledhill/robojack/web/oauth"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	state := oauth.NewState()
	c.Redirect(http.StatusTemporaryRedirect, oauth.Config.AuthCodeURL(state))
}

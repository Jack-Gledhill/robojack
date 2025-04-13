package dashboard

import (
	"net/http"

	"github.com/Jack-Gledhill/robojack/bot"
	"github.com/Jack-Gledhill/robojack/debug"
	"github.com/Jack-Gledhill/robojack/web/oauth"
	"github.com/Jack-Gledhill/robojack/web/src/templates"
	"github.com/Jack-Gledhill/robojack/web/utils"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	claims := utils.GetClaimsFromCtx(c)
	botUser := &oauth.User{
		AvatarHash:  bot.Session.State.User.Avatar,
		DisplayName: bot.Session.State.User.GlobalName,
		ID:          bot.Session.State.User.ID,
		Username:    bot.Session.State.User.Username,
	}

	c.HTML(http.StatusOK, "", templates.Index(botUser, claims.User, debug.System, debug.Build, debug.Git, debug.Runtime.Snapshot()))
}

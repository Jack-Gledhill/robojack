package auth

import (
	"context"
	"net/http"

	"github.com/Jack-Gledhill/robojack/config"
	"github.com/Jack-Gledhill/robojack/log"
	"github.com/Jack-Gledhill/robojack/web/jwt"
	"github.com/Jack-Gledhill/robojack/web/oauth"
	"github.com/Jack-Gledhill/robojack/web/response"

	"github.com/gin-gonic/gin"
)

var jwtCookieMaxAge = jwt.Validity.Seconds()

func callback(c *gin.Context) {
	// Check for a valid state
	if !oauth.PopState(c.Query("state")) {
		response.New().
			Status(http.StatusBadRequest).
			Msg("invalid oauth state").
			Send(c)
		return
	}

	// Fetch code from query string or fail if not given
	code := c.Query("code")
	if code == "" {
		response.New().
			Status(http.StatusBadRequest).
			Msg("no code provided").
			Send(c)
		return
	}

	// Exchange code for bearer token
	token, err := oauth.Config.Exchange(context.Background(), code)
	if err != nil {
		log.Error("Failed to exchange oauth code for token: %s", err.Error())

		response.New().
			Status(http.StatusInternalServerError).
			Msg("failed to exchange code for token").
			Send(c)
		return
	}

	// Use bearer token to fetch user details
	user, err := oauth.GetUser(token)
	if err != nil {
		log.Error("Failed to fetch user data: %s", err.Error())

		response.New().
			Status(http.StatusInternalServerError).
			Msg("failed to fetch user data").
			Send(c)
		return
	}

	// Check that authenticated user has access rights
	if user.ID != config.Bot.OwnerID {
		response.New().
			Status(http.StatusForbidden).
			Msg("user does not have access rights").
			Send(c)
		return
	}

	// Create a new JWT for the user
	userToken, err := jwt.New(user)
	if err != nil {
		log.Error("Failed to create JWT for user %s (%d): %s", user.Username, user.ID, err.Error())

		response.New().
			Status(http.StatusInternalServerError).
			Msg("failed to create new JWT").
			Send(c)
		return
	}

	// Set the JWT as a cookie on the user's browser
	c.SetCookie(config.Web.JWT.Cookie, userToken, int(jwtCookieMaxAge), "/", config.Web.Domain().Hostname(), false, false)

	// All being well, redirect to the dashboard
	c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
}

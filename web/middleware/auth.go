package middleware

import (
	"github.com/Jack-Gledhill/robojack/config"
	"net/http"

	"github.com/Jack-Gledhill/robojack/log"
	"github.com/Jack-Gledhill/robojack/web/jwt"
	"github.com/Jack-Gledhill/robojack/web/response"

	"github.com/gin-gonic/gin"
)

// Authentication checks for a JWT in cookies, then runs validation checks and passes the claims to the handler.
// If any of the prior tasks fail, the middleware will abort the request and show an error to the client.
func Authentication(c *gin.Context) {
	// Try fetch the token from cookies
	token, err := c.Cookie(config.Web.JWT.Cookie)
	if err != nil {
		// Instead try to fetch from headers
		token = c.Request.Header.Get("authorization")
		if token == "" {
			response.New().
				Status(http.StatusUnauthorized).
				Msg("not authorized").
				Send(c)

			c.Abort()
			return
		}
	}

	// Run validation checks on token
	valid, claims, err := jwt.Validate(token)
	if err != nil {
		log.Error("Error when validating JWT: %s", err.Error())

		response.New().
			Status(http.StatusInternalServerError).
			Msg("error when validating authorization token").
			Send(c)
		c.Abort()
		return
	}

	// Error if JWT is invalid
	if !valid {
		response.New().
			Status(http.StatusUnauthorized).
			Msg("invalid or expired token").
			Send(c)
		c.Abort()
		return
	}

	c.Set("claims", claims)
	c.Next()
}

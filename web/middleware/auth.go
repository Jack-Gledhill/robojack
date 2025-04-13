package middleware

import (
	"net/http"

	"github.com/Jack-Gledhill/robojack/config"
	"github.com/Jack-Gledhill/robojack/web/jwt"
	"github.com/Jack-Gledhill/robojack/web/response"

	"github.com/gin-gonic/gin"
)

func parseAuth(c *gin.Context) (bool, *jwt.Claims) {
	// Try fetch the token from cookies
	token, err := c.Cookie(config.Web.JWT.Cookie)
	if err != nil {
		// Instead try to fetch from headers
		token = c.Request.Header.Get("authorization")
		if token == "" {
			return false, nil
		}
	}

	// Run validation checks on token
	valid, claims, err := jwt.Validate(token)
	if err != nil {
		return false, nil
	}

	return valid, claims
}

// Authentication checks for a JWT in cookies, then runs validation checks and passes the claims to the handler.
// If any of the prior tasks fail, the middleware will abort the request and show an error to the client.
func Authentication(c *gin.Context) {
	valid, claims := parseAuth(c)
	if !valid {
		response.New().
			Status(http.StatusUnauthorized).
			Msg("not authorized").
			Send(c)
		c.Abort()
		return
	}

	c.Set("claims", claims)
	c.Next()
}

// AuthenticationWithRedirect behaves similarly to Authentication, but will redirect the client to the login route if they're not authenticated
func AuthenticationWithRedirect(c *gin.Context) {
	valid, claims := parseAuth(c)
	if !valid {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
		c.Abort()

		return
	}

	c.Set("claims", claims)
	c.Next()
}

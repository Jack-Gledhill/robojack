package auth

import (
	"net/http"

	"github.com/Jack-Gledhill/robojack/log"
	"github.com/Jack-Gledhill/robojack/web/jwt"
	"github.com/Jack-Gledhill/robojack/web/response"

	"github.com/gin-gonic/gin"
)

func me(c *gin.Context) {
	token, err := c.Cookie(cookieToken)
	if err != nil {
		response.New().
			Status(http.StatusUnauthorized).
			Msg("not authorized").
			Send(c)
		return
	}

	valid, claims, err := jwt.Validate(token)
	if err != nil {
		log.Error("Error when validating JWT: %s", err.Error())

		response.New().
			Status(http.StatusInternalServerError).
			Msg("error when validating authorization token").
			Send(c)
		return
	}

	if !valid {
		response.New().
			Status(http.StatusUnauthorized).
			Msg("invalid or expired token").
			Send(c)
		return
	}

	response.New().
		Data(claims.User).
		Send(c)
}

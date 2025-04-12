package auth

import (
	"github.com/Jack-Gledhill/robojack/web/response"
	"github.com/Jack-Gledhill/robojack/web/utils"

	"github.com/gin-gonic/gin"
)

func me(c *gin.Context) {
	claims := utils.GetClaimsFromCtx(c)
	response.New().
		Data(claims.User).
		Send(c)
}

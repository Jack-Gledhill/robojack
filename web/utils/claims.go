package utils

import (
	"github.com/Jack-Gledhill/robojack/web/jwt"

	"github.com/gin-gonic/gin"
)

// GetClaimsFromCtx takes in a gin.Context and fetches the JWT claims from it.
// Claims will only be present in a context when added by middleware.Authentication
// If no claims are present in the context, nil is returned instead
func GetClaimsFromCtx(c *gin.Context) *jwt.Claims {
	claims, ok := c.Get("claims")
	if !ok {
		return nil
	}

	return claims.(*jwt.Claims)
}

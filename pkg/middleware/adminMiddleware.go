package mw

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jwtHelper "github.com/horzu/golang/cart-api/pkg/jwt"
)

func AdminAuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil && decodedClaims.Role == "admin" {
				c.Set("userID", decodedClaims.UserId)
				c.Next()
				c.Abort()
				return
			}

			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized"})
			c.Abort()
			return
		}
	}
}



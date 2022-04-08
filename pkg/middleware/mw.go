package mw

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jwtHelper "github.com/horzu/golang/cart-api/pkg/jwt"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil && decodedClaims.IsAdmin {
				c.Next()
				c.Abort()
				return
			}

			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		}
	}
}

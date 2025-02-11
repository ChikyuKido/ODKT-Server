package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"odkt/server/store"
	"odkt/server/util"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authenticationHeader := c.GetHeader("Authentication")
		if authenticationHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no authentication header provided"})
			return
		}
		if !strings.HasPrefix(authenticationHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authentication header in wrong format. Has to start with Bearer"})
			return
		}
		tokenString := strings.TrimPrefix(authenticationHeader, "Bearer ")
		token, err := util.GetToken(tokenString)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		uuid, ok := claims["uuid"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		}
		var user = store.GetUserByUUID(uuid)
		if user == nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "something went wrong. Please try again later."})
		}
		c.Set("user", user)
	}
}

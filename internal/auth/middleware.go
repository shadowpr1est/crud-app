package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const UserCtxKey = "user"

func (j *JWTManager) GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		parts := strings.Split(header, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid header"})
			return
		}
		claims, err := j.VerifyToken(parts[1])

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set(UserCtxKey, claims)
		c.Next()
	}
}

func GetUser(c *gin.Context) (*UserClaims, bool) {
	v, ok := c.Get(UserCtxKey)
	if !ok {
		return nil, false
	}
	claims, ok := v.(*UserClaims)
	return claims, ok
}

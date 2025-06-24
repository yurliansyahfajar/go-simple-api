package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/yurliansyahfajar/go-simple-api/utils"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{"message": "unauthorize user"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": "unauthorize user"})
		return
	}

	c.Set("userId", userId)
	c.Next()
}

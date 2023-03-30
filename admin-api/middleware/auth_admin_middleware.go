package middleware

import (
	"github.com/gin-gonic/gin"
	jwt_helper "github.com/sx931210/app-mall/utils/jwt"
	"net/http"
)

func AuthAdminMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			authorization := c.GetHeader("Authorization")
			decodedClaims, err := jwt_helper.VerifyToken(authorization, secretKey)
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"message": "身份验证失败"})
				c.Abort()
				return
			}
			c.Set("uid", decodedClaims.Id)
			c.Next()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "未授权"})
			c.Abort()
			return
		}
	}
}

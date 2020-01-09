package middleware

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/auth"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tk, valid := auth.ValidateToken(c.Request)
		if !valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user_id", tk.UserID)
		c.Set("email", tk.Email)
		c.Next()
	}
}

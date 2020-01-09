package util

import "github.com/gin-gonic/gin"

func GetUserIDFromContext(c *gin.Context) (uint, bool) {
	userID, ok := c.MustGet("user_id").(uint)
	return userID, ok
}

func GetEmailFromContext(c *gin.Context) (string, bool) {
	email, ok := c.MustGet("email").(string)
	return email, ok
}
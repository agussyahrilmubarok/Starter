package helper

import "github.com/gin-gonic/gin"

func GetUserIDFromGin(c *gin.Context) string {
	value, exists := c.Get("user_id")
	if !exists {
		return ""
	}

	userID, ok := value.(string)
	if !ok {
		return ""
	}

	return userID
}

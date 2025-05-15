package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseQueryInt(c *gin.Context, key string, defaultValue int) int {
	if val := c.Query(key); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			return parsed
		}
	}
	return defaultValue
}

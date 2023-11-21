package util

import (
	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, code int, message string, keyData string, data interface{}) {
	c.JSON(code, gin.H{
		"success": true,
		"message": message,
		keyData:   data,
	})
}

func FailOrErrorResponse(c *gin.Context, code int, message string, detail interface{}) {
	c.JSON(code, gin.H{
		"success": false,
		"message": message,
		"data":    detail,
	})
}

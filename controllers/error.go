package controllers

import "github.com/gin-gonic/gin"

func defaultErrResponse(c *gin.Context, statusCode, errorCode int, errorMsg string) {
	c.JSON(statusCode, gin.H{
		"error":     errorMsg,
		"errorCode": errorCode,
	})
}

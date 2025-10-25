package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getHealth(c *gin.Context) {
	isHealthy, err := isHealthy()
	if isHealthy {
		c.IndentedJSON(http.StatusOK, gin.H{
			"isHealthy": isHealthy,
			"TimeStamp": time.Now().Format(time.DateTime),
		})
	} else {
		c.IndentedJSON(http.StatusServiceUnavailable, gin.H{
			"isHealthy": isHealthy,
			"TimeStamp": time.Now().Format(time.DateTime),
			"error":     err.Error(),
		})
	}
}

package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func handleError(err error, c *gin.Context) {
	if err == nil {
		return
	}
	log.Error(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"message":   err.Error(),
		"code":      http.StatusInternalServerError,
		"timestamp": time.Now(),
	})
}

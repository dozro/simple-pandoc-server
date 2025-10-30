package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func handleError(err error, c *gin.Context) bool {
	if err == nil {
		return false
	}
	log.Error(err)
	errn := c.AbortWithError(http.StatusInternalServerError, err)
	log.Fatal(errn)
	return true
}

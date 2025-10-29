package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func extractDataFromReq(c *gin.Context) ([]byte, error) {
	log.Debugf("trying to extract data from %s via %s", c.Request.Host, c.Request.URL.String())
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("File upload error: %s", err.Error()))
		log.Errorf("File upload error: %s", c.Request.URL.String())
		return nil, err
	}
	f, err := file.Open()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("File upload error: %s (File could not be opened)", c.Request.URL.String()))
		log.Errorf("File upload error: %s", c.Request.URL.String())
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("File upload error: %s", c.Request.URL.String()))
		log.Errorf("File upload error: %s", c.Request.URL.String())
		return nil, err
	}
	return data, nil
}

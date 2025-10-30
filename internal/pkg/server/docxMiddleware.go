package server

import (
	"context"
	"simple-pandoc-server/internal/pkg/convert"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ParseDocxRawToPDF(c *gin.Context) {
	log.Debugf("trying to parse docx data received from %s", c.Request.Host)
	data, err := extractDataFromReq(c)
	handleError(err, c)
	out, err := concurrentCacheLookupAndRendering(context.Background(), data, convert.ParseDocxToHTML)
	handleError(err, c)
	c.Data(200, "application/pdf", out)
}

func ParseDocxRawToHTML(c *gin.Context) {
	log.Debugf("trying to parse docx data received from %s", c.Request.Host)
	data, err := extractDataFromReq(c)
	handleError(err, c)
	out, err := concurrentCacheLookupAndRendering(context.Background(), data, convert.ParseDocxToHTML)
	handleError(err, c)
	c.Data(200, "text/html", out)
}

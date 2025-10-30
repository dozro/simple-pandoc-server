package server

import (
	"context"
	"simple-pandoc-server/internal/pkg/convert"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ParseLatexRawToPDF(c *gin.Context) {
	log.Debugf("trying to parse latex data received from %s", c.Request.Host)
	data, err := extractDataFromReq(c)
	if handleError(err, c) {
		return
	}
	out, err := concurrentCacheLookupAndRendering(context.Background(), data, convert.ParseLatexDataToPdf)
	if handleError(err, c) {
		return
	}
	c.Data(200, "application/pdf", out)
}

func ParseLatexRawToHTML(c *gin.Context) {
	log.Debugf("trying to parse latex data received from %s", c.Request.Host)
	data, err := extractDataFromReq(c)
	if handleError(err, c) {
		return
	}
	out, err := concurrentCacheLookupAndRendering(context.Background(), data, convert.ParseLatexDataToHtml)
	if handleError(err, c) {
		return
	}
	c.Data(200, "text/html", out)
}

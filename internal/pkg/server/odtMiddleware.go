package server

import (
	"context"
	"simple-pandoc-server/internal/pkg/convert"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ParseOdtRawToHTML(c *gin.Context) {
	log.Debugf("trying to parse odt string from %s via %s", c.Request.Host, c.Request.URL.String())
	data, err := extractDataFromReq(c)
	if handleError(err, c) {
		return
	}
	out, err := concurrentCacheLookupAndRendering(context.Background(), data, convert.ParseOdtDataToHtml)
	if handleError(err, c) {
		return
	}
	c.Data(200, "text/html", out)
}

func ParseOdtRawToPdf(c *gin.Context) {
	log.Debugf("trying to parse odt string from %s via %s", c.Request.Host, c.Request.URL.String())
	data, err := extractDataFromReq(c)
	if handleError(err, c) {
		return
	}
	out, err := concurrentCacheLookupAndRendering(context.Background(), data, convert.ParseOdtDataToPdf)
	if handleError(err, c) {
		return
	}
	c.Data(200, "application/pdf", out)
}

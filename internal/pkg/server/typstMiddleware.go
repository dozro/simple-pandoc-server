package server

import (
	"context"
	"simple-pandoc-server/internal/pkg/convert"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ParseTypstRawToHtml(c *gin.Context) {
	log.Debugf("trying to parse typst string from %s via %s", c.Request.Host, c.Request.URL.String())
	data, err := extractDataFromReq(c)
	if handleError(err, c) {
		return
	}
	out, err := concurrentCacheLookupAndRendering(context.Background(), data, convert.ParseTypstDataToHtml)
	if handleError(err, c) {
		return
	}
	c.Data(200, "text/html", out)
}

func ParseTypstRawToPdf(c *gin.Context) {
	log.Debugf("trying to parse typst string from %s via %s", c.Request.Host, c.Request.URL.String())
	data, err := extractDataFromReq(c)
	if handleError(err, c) {
		return
	}
	out, err := concurrentCacheLookupAndRendering(context.Background(), data, convert.ParseTypstDataToPdf)
	if handleError(err, c) {
		return
	}
	c.Data(200, "application/pdf", out)
}

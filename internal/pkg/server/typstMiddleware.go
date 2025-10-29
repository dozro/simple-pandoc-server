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
	handleError(err, c)
	out, err := concurrentCacheLookupAndRendering(context.Background(), data, convert.ParseTypstDataToHtml)
	handleError(err, c)
	c.Data(200, "text/html", out)
}

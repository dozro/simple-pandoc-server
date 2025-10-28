package convert

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TypstData struct {
	TypstString string `json:"typstString"`
}

func ParseTypstRawToHtml(c *gin.Context) {
	log.Debugf("trying to parse typst string from %s via %s", c.Request.Host, c.Request.URL.String())
	data, err := extractDataFromReq(c)
	if err != nil {
		log.Error(err)
	}
	out, _ := convertToHtmlUsingPandoc("typst", data)
	c.Data(200, "text/html", out.Bytes())
}

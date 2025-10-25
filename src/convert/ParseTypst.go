package convert

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TypstData struct {
	typstString string `json:"typstString"`
}

func ParseTypstRawToHtml(c *gin.Context) {
	log.Debugf("trying to parse typst string from %s via %s", c.Request.Host, c.Request.URL.String())
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("File upload error: %s", err.Error()))
		log.Errorf("File upload error: %s", c.Request.URL.String())
	}
	f, err := file.Open()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("File upload error: %s (File could not be opened)", c.Request.URL.String()))
		log.Errorf("File upload error: %s", c.Request.URL.String())
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("File upload error: %s", c.Request.URL.String()))
		log.Errorf("File upload error: %s", c.Request.URL.String())
	}
	out, _ := convertToHtmlUsingPandoc("typst", data)
	c.Data(200, "text/html", out.Bytes())
}

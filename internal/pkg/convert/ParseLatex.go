package convert

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rwestlund/gotex"
	log "github.com/sirupsen/logrus"
)

type LatexData struct {
	LatexString string `json:"latexString"`
}

func ParseLatexRawToPDF(c *gin.Context) {
	if isGoTexEnabled {
		parseLatexUsingGoTexRawToPDF(c)
	} else {
		parseLatexUsingPandocRawToPdf(c)
	}
}

func parseLatexUsingGoTexRawToPDF(c *gin.Context) {
	log.Debugf("trying to parse latex data received from %s via %s", c.Request.Host, c.Request.URL.String())
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
	pdf, err := gotex.Render(string(data), gotex.Options{
		Command: os.Getenv("LATEX_COMMAND"),
		Runs:    1,
	})
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("failed to render file: %s", c.Request.URL.String()))
		log.Errorf("failed to render file: %s (error: %s)", c.Request.URL.String(), err.Error())
	}
	_ = f.Close()
	c.Data(200, "application/pdf", []byte(pdf))
	log.Debugf("successfully parsed latex data from %s via %s", c.Request.Host, c.Request.URL.String())
}

func ParseLatexPlainToPdf(c *gin.Context) {
	if isGoTexEnabled {
		parseLatexUsingGoTexPlainToPdf(c)
	} else {
		parseLatexUsingPandocPlainToPdf(c)
	}
}

func parseLatexUsingPandocPlainToPdf(c *gin.Context) {
	log.Debugf("trying to parse latex string from %s via %s", c.Request.Host, c.Request.URL.String())
	var data LatexData
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	out, _ := convertToPdfUsingPandoc("latex", []byte(data.LatexString))
	c.Data(200, "application/pdf", out.Bytes())
}

func parseLatexUsingGoTexPlainToPdf(c *gin.Context) {
	log.Debugf("trying to parse latex string from %s via %s", c.Request.Host, c.Request.URL.String())
	var data LatexData
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	pdf, err := gotex.Render(data.LatexString, gotex.Options{
		Command: os.Getenv("LATEX_COMMAND"),
		Runs:    1,
	})
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("failed to render file: %s", c.Request.URL.String()))
		log.Errorf("failed to render file: %s (error: %s)", c.Request.URL.String(), err.Error())
	}
	c.Data(200, "application/pdf", []byte(pdf))
	log.Debugf("successfully parsed latex data from %s via %s", c.Request.Host, c.Request.URL.String())
}

func ParseLatexPlainToHtml(c *gin.Context) {
	log.Debugf("trying to parse latex string from %s via %s", c.Request.Host, c.Request.URL.String())
	var data LatexData
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	out, _ := convertToHtmlUsingPandoc("latex", []byte(data.LatexString))
	c.Data(200, "text/html", out.Bytes())
}

func parseLatexUsingPandocRawToPdf(c *gin.Context) {
	log.Debugf("trying to parse latex string from %s via %s", c.Request.Host, c.Request.URL.String())
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
	out, _ := convertToPdfUsingPandoc("latex", data)
	c.Data(200, "application/pdf", out.Bytes())
}

func ParseLatexRawToHtml(c *gin.Context) {
	log.Debugf("trying to parse latex string from %s via %s", c.Request.Host, c.Request.URL.String())
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
	out, _ := convertToHtmlUsingPandoc("latex", data)
	c.Data(200, "text/html", out.Bytes())
}

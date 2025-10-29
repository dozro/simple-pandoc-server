package convert

import (
	"bytes"

	log "github.com/sirupsen/logrus"
)

type TypstData struct {
	TypstString string `json:"typstString"`
}

func ParseTypstDataToHtml(d []byte) (bytes.Buffer, error) {
	log.Debugf("starting conversion of typst data to html")
	out, err := convertToHtmlUsingPandoc("typst", d)
	if err != nil {
		return bytes.Buffer{}, err
	}
	return out, nil
}

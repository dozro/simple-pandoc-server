package convert

import (
	log "github.com/sirupsen/logrus"
)

type TypstData struct {
	TypstString string `json:"typstString"`
}

func ParseTypstDataToHtml(d []byte) ([]byte, error) {
	log.Debugf("starting conversion of typst data to html")
	out, err := convertToHtmlUsingPandoc("typst", d)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

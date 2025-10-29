package convert

import log "github.com/sirupsen/logrus"

func ParseOdtDataToHtml(d []byte) ([]byte, error) {
	log.Debugf("starting conversion of typst data to html")
	out, err := convertToHtmlUsingPandoc("odt", d)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func ParseOdtDataToPdf(d []byte) ([]byte, error) {
	log.Debugf("starting conversion of typst data to html")
	out, err := convertToPdfUsingPandoc("odt", d)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

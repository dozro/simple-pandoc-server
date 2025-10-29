package convert

import (
	"os"

	"github.com/rwestlund/gotex"
)

func ParseLatexDataToPdf(d []byte) ([]byte, error) {
	if isGoTexEnabled {
		return parseLatexUsingGoTexRawToPDF(d)
	} else {
		return parseLatexUsingPandocRawToPdf(d)
	}
}

func parseLatexUsingGoTexRawToPDF(data []byte) ([]byte, error) {
	pdf, err := gotex.Render(string(data), gotex.Options{
		Command: os.Getenv("LATEX_COMMAND"),
		Runs:    1,
	})
	if err != nil {
		return nil, err
	}
	return pdf, nil
}

func parseLatexUsingPandocRawToPdf(d []byte) ([]byte, error) {
	out, err := convertToPdfUsingPandoc("latex", d)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func ParseLatexRawToHtml(d []byte) ([]byte, error) {
	out, err := convertToHtmlUsingPandoc("latex", d)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

package convert

import (
	"os"

	"github.com/rwestlund/gotex"
)

func ParseLatexDataToPdf(d []byte) ([]byte, error) {
	if isGoTexEnabled {
		return parseLatexUsingGoTexDataToPDF(d)
	} else {
		return parseLatexUsingPandocDataToPdf(d)
	}
}

func parseLatexUsingGoTexDataToPDF(data []byte) ([]byte, error) {
	pdf, err := gotex.Render(string(data), gotex.Options{
		Command: os.Getenv("LATEX_COMMAND"),
		Runs:    1,
	})
	if err != nil {
		return nil, err
	}
	return pdf, nil
}

func parseLatexUsingPandocDataToPdf(d []byte) ([]byte, error) {
	out, err := convertToPdfUsingPandoc("latex", d)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func ParseLatexDataToHtml(d []byte) ([]byte, error) {
	out, err := convertToHtmlUsingPandoc("latex", d)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

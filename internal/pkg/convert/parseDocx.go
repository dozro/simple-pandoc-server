package convert

func ParseDocxToHTML(data []byte) ([]byte, error) {
	out, err := convertToHtmlUsingPandoc("docx", data)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func ParseDocxToPdf(data []byte) ([]byte, error) {
	out, err := convertToPdfUsingPandoc("docx", data)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

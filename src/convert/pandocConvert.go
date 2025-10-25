package convert

import (
	"bytes"
	"context"
	"os"
	"os/exec"
)

func convertToHtmlUsingPandoc(sourceFormat string, doc []byte) (bytes.Buffer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, os.Getenv("PANDOC_COMMAND"), "-f", "latex", "-t", "html")
	cmd.Stdin = bytes.NewReader(doc)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out, err
}

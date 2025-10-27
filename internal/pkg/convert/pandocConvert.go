package convert

import (
	"bytes"
	"context"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func convertToHtmlUsingPandoc(sourceFormat string, doc []byte) (bytes.Buffer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, os.Getenv("PANDOC_COMMAND"), "-f", sourceFormat, "-t", "html")
	cmd.Stdin = bytes.NewReader(doc)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		log.Errorf("Pandoc command timed out")
		return out, ctx.Err()
	}
	return out, err
}

func convertToPdfUsingPandoc(sourceFormat string, doc []byte) (bytes.Buffer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, os.Getenv("PANDOC_COMMAND"), "-f", sourceFormat, "-t", "pdf")
	cmd.Stdin = bytes.NewReader(doc)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		log.Errorf("Pandoc command timed out")
		return out, ctx.Err()
	}
	return out, err
}

package convert

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

func generateMathRenderingCmdFlag() string {
	if len(mathRenderingConfig.MathRenderingURL) == 0 && (mathRenderingConfig.MathRenderingEngine == Mathjax || mathRenderingConfig.MathRenderingEngine == Webtex || mathRenderingConfig.MathRenderingEngine == Katex) {
		return fmt.Sprintf("--%s=%s", MathRenderingEngineName[mathRenderingConfig.MathRenderingEngine], mathRenderingConfig.MathRenderingURL)
	} else {
		return fmt.Sprintf("--%s", MathRenderingEngineName[mathRenderingConfig.MathRenderingEngine])
	}
}

func convertToHtmlUsingPandoc(sourceFormat string, doc []byte) (bytes.Buffer, error) {
	log.Debugf("Converting %s document to html with pandoc", sourceFormat)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, os.Getenv("PANDOC_COMMAND"), "-f", sourceFormat, "-t", "html", generateMathRenderingCmdFlag())
	log.Debugf("Executing: %s", strings.Join(cmd.Args, " "))
	cmd.Stdin = bytes.NewReader(doc)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		log.Errorf("Pandoc command timed out")
		return out, ctx.Err()
	}
	log.Debugf("Pandoc command finished executing will return the result: %s", out.String())
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

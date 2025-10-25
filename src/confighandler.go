package main

import (
	"flag"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Timeout       time.Duration
	Debug         bool
	ListenOnIP    string
	LatexCommand  string
	UseGoTex      bool
	PandocCommand string
	typstCommand  string
	trustedProxy  string
}

func readConfigFromEnv() Config {
	debugOutFlag := flag.Bool("debug", os.Getenv("DEBUG") == "true", "Enable debug mode")
	var defaultTimeout time.Duration
	if len(os.Getenv("TIMEOUT")) != 0 {
		timeoutEnv, err := strconv.Atoi(os.Getenv("TIMEOUT"))
		if err != nil {
			defaultTimeout = 30 * time.Second
			log.Debugf("Error parsing TIMEOUT environment variable: %v (is it a valid int?)", err)
		}
		defaultTimeout = time.Duration(timeoutEnv) * time.Second
	} else {
		defaultTimeout = 30 * time.Second
	}
	timeoutFlag := flag.Duration("timeout", defaultTimeout, "Timeout for rendering")
	listenOnIpFlag := flag.String("listen-on", os.Getenv("LISTEN_ON"), "Listen on IP address Format: \"ip:port\"")
	latexCommandFlag := flag.String("latex-command", os.Getenv("LATEX_COMMAND"), "the path to pdflatex or equiv")
	pandocCommandFlag := flag.String("pandoc-command", os.Getenv("PANDOC_COMMAND"), "the path to pandoc")
	typstCommandFlag := flag.String("typst-command", os.Getenv("TYPST_COMMAND"), "the path to typst")
	defaultGoTexEnable := true
	if os.Getenv("GOTEX_ENABLE") == "false" {
		defaultGoTexEnable = false
	}
	enableGoTexFlag := flag.Bool("enable-gotex", defaultGoTexEnable, "Enable GoTex compatibility, if disabled it will use pandoc")
	var defaultTrustedProxy string
	if len(os.Getenv("TRUSTED_PROXY")) != 0 {
		defaultTrustedProxy = os.Getenv("TRUSTED_PROXY")
	} else {
		defaultTrustedProxy = "0.0.0.0"
	}
	trustedProxyFlag := flag.String("trust-proxy", defaultTrustedProxy, "the ip of the proxy server")
	flag.Parse()
	return Config{
		Timeout:       *timeoutFlag,
		Debug:         *debugOutFlag,
		ListenOnIP:    *listenOnIpFlag,
		LatexCommand:  *latexCommandFlag,
		PandocCommand: *pandocCommandFlag,
		typstCommand:  *typstCommandFlag,
		trustedProxy:  *trustedProxyFlag,
		UseGoTex:      *enableGoTexFlag,
	}
}

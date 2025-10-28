package confighandling

import (
	"flag"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Timeout             time.Duration `json:"timeout"`
	Debug               bool          `json:"debug"`
	ListenOnIP          string        `json:"listen_on_ip"`
	LatexCommand        string        `json:"latex_command"`
	UseGoTex            bool          `json:"use_gotex"`
	MathRenderingEngine string        `json:"math_rendering_engine"`
	MathRenderingURL    string        `json:"math_rendering_url"`
	PandocCommand       string        `json:"pandoc_command"`
	TypstCommand        string        `json:"typst_command"`
	TrustedProxy        string        `json:"trusted_proxy"`
	PresharedKey        string        `json:"preshared_key"`
}

func ReadConfigFromEnv() Config {
	// whether debug mode should be enabled
	debugOutFlag := flag.Bool("debug", os.Getenv("DEBUG") == "true", "Enable debug mode")
	// command timeout
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
	// ip to listen on
	listenOnIpFlag := flag.String("listen-on", os.Getenv("LISTEN_ON"), "Listen on IP address Format: \"ip:port\"")
	// latex command for gotex
	latexCommandFlag := flag.String("latex-command", os.Getenv("LATEX_COMMAND"), "the path to pdflatex or equiv")
	// pandoc command
	pandocCommandFlag := flag.String("pandoc-command", os.Getenv("PANDOC_COMMAND"), "the path to pandoc")
	// typst command
	typstCommandFlag := flag.String("typst-command", os.Getenv("TYPST_COMMAND"), "the path to typst")
	// whether go tex is enabled
	defaultGoTexEnable := true
	if os.Getenv("GOTEX_ENABLE") == "false" {
		defaultGoTexEnable = false
	}
	enableGoTexFlag := flag.Bool("enable-gotex", defaultGoTexEnable, "Enable GoTex compatibility, if disabled it will use pandoc")
	// trusted proxy
	var defaultTrustedProxy string
	if len(os.Getenv("TRUSTED_PROXY")) != 0 {
		defaultTrustedProxy = os.Getenv("TRUSTED_PROXY")
	} else {
		defaultTrustedProxy = "0.0.0.0"
	}
	trustedProxyFlag := flag.String("trust-proxy", defaultTrustedProxy, "the ip of the proxy server")
	// math rendering engine
	var defaultMathRenderingEngine string
	if len(os.Getenv("MATH_RENDERING_ENGINE")) != 0 {
		defaultMathRenderingEngine = os.Getenv("MATH_RENDERING_ENGINE")
	} else {
		defaultMathRenderingEngine = "mathml"
	}
	mathRenderingEngineFlag := flag.String("math-rendering-engine", defaultMathRenderingEngine, "The math rendering engine to use")
	mathRenderingURLFlag := flag.String("math-rendering-url", os.Getenv("MATH_RENDERING_URL"), "The url of the math rendering server")
	flag.Parse()
	return Config{
		Timeout:             *timeoutFlag,
		Debug:               *debugOutFlag,
		ListenOnIP:          *listenOnIpFlag,
		LatexCommand:        *latexCommandFlag,
		PandocCommand:       *pandocCommandFlag,
		TypstCommand:        *typstCommandFlag,
		TrustedProxy:        *trustedProxyFlag,
		UseGoTex:            *enableGoTexFlag,
		MathRenderingEngine: *mathRenderingEngineFlag,
		MathRenderingURL:    *mathRenderingURLFlag,
	}
}

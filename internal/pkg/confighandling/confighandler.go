package confighandling

import (
	"flag"
	"os"
	"time"
)

type Config struct {
	Timeout              time.Duration `json:"timeout"`
	Debug                bool          `json:"debug"`
	ListenOnIP           string        `json:"listen_on_ip"`
	LatexCommand         string        `json:"latex_command"`
	UseGoTex             bool          `json:"use_gotex"`
	MathRenderingEngine  string        `json:"math_rendering_engine"`
	MathRenderingURL     string        `json:"math_rendering_url"`
	PandocCommand        string        `json:"pandoc_command"`
	TypstCommand         string        `json:"typst_command"`
	TrustedProxy         string        `json:"trusted_proxy"`
	PresharedKey         string        `json:"preshared_key"`
	CacheCleanupInterval time.Duration `json:"cache_cleanup_interval"`
	CacheExpiration      time.Duration `json:"cache_expiration"`
}

func ReadConfigFromEnv() Config {
	// whether debug mode should be enabled
	debugOutFlag := flag.Bool("debug", defValBool("DEBUG", false), "Enable debug mode")
	// command timeout
	timeoutFlag := flag.Duration("timeout", defValTimeDuration("TIMEOUT", 30*time.Second, time.Second), "Timeout for rendering")
	// ip to listen on
	listenOnIpFlag := flag.String("listen-on", defValString("LISTEN_ON", "0.0.0.0:3030"), "Listen on IP address Format: \"ip:port\"")
	// latex command for gotex
	latexCommandFlag := flag.String("latex-command", os.Getenv("LATEX_COMMAND"), "the path to pdflatex or equiv")
	// pandoc command
	pandocCommandFlag := flag.String("pandoc-command", os.Getenv("PANDOC_COMMAND"), "the path to pandoc")
	// typst command
	typstCommandFlag := flag.String("typst-command", os.Getenv("TYPST_COMMAND"), "the path to typst")
	// whether go tex is enabled
	enableGoTexFlag := flag.Bool("enable-gotex", defValBool("GOTEX_ENABLE", false), "Enable GoTex compatibility, if disabled it will use pandoc")
	// trusted proxy
	trustedProxyFlag := flag.String("trust-proxy", defValString("TRUSTED_PROXY", "0.0.0.0"), "the ip of the proxy server")
	// math rendering engine
	mathRenderingEngineFlag := flag.String("math-rendering-engine", defValString("MATH_RENDERING_ENGINE", "mathml"), "The math rendering engine to use")
	mathRenderingURLFlag := flag.String("math-rendering-url", os.Getenv("MATH_RENDERING_URL"), "The url of the math rendering server")
	cacheCleanupIntervalFlag := flag.Duration("cache-cleanup-interval", defValTimeDuration("CACHE_CLEANUP_INTERVAL", 10*time.Minute, time.Minute), "Cache cleanup interval (in Minutes)")
	cacheExpirationFlag := flag.Duration("cache-expire-after", defValTimeDuration("CACHE_EXPIRATION", 5*time.Minute, time.Minute), "Cache expiration after x minutes")
	flag.Parse()
	return Config{
		Timeout:              *timeoutFlag,
		Debug:                *debugOutFlag,
		ListenOnIP:           *listenOnIpFlag,
		LatexCommand:         *latexCommandFlag,
		PandocCommand:        *pandocCommandFlag,
		TypstCommand:         *typstCommandFlag,
		TrustedProxy:         *trustedProxyFlag,
		UseGoTex:             *enableGoTexFlag,
		MathRenderingEngine:  *mathRenderingEngineFlag,
		MathRenderingURL:     *mathRenderingURLFlag,
		CacheCleanupInterval: *cacheCleanupIntervalFlag,
		CacheExpiration:      *cacheExpirationFlag,
	}
}

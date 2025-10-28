package checks

import (
	"bytes"
	"os"
	"os/exec"
	cfgh "simple-pandoc-server/internal/pkg/confighandling"
	"simple-pandoc-server/internal/pkg/convert"
	"simple-pandoc-server/internal/pkg/zero"

	log "github.com/sirupsen/logrus"
)

func searchPackageAndSetEnv(pkgname string, envname string) {
	path, err := exec.LookPath(pkgname)
	if err != nil {
		log.Errorf("Could not find package %s in path not setting it as Env %s, it has to manually be set or cmdline args have to be used: %s", pkgname, envname, err)
		return
	}
	log.Infof("found package %s in path setting it as Env %s: %s", pkgname, envname, path)
	_ = os.Setenv(envname, path)
}

func determineVersionOfPackageAndSetEnv(command string, envName string) {
	versioncmd := exec.Command(command, "--version")
	headcmd := exec.Command("head", "-n", "1")
	var out bytes.Buffer
	headcmd.Stdin, _ = versioncmd.StdoutPipe()
	headcmd.Stdout = &out
	err := versioncmd.Start()
	err = headcmd.Start()
	err = versioncmd.Wait()
	if err != nil {
		log.Errorf("Could not determine version of package %s (does it exist and is supported?): %s", command, err)
		return
	}
	version := out.String()
	log.Infof("determining version of package %s and setting in env %s: %s", command, envName, version)
	os.Setenv(envName, version)
}

func PreflightPackageSearch() {
	log.Info("searching for packages")
	go searchPackageAndSetEnv("pandoc", "PANDOC_COMMAND")
	go searchPackageAndSetEnv("pdflatex", "LATEX_COMMAND")
	go searchPackageAndSetEnv("typst", "TYPST_COMMAND")
}

func setMathRenderingEngine(config cfgh.Config) {
	log.Info("setting math rendering engine for pandoc")
	var mathrenderingengine convert.MathRenderingEngine
	switch config.MathRenderingEngine {
	case "mathjax":
		mathrenderingengine = convert.Mathjax
	case "mathml":
		mathrenderingengine = convert.Mathml
	case "webtex":
		mathrenderingengine = convert.Webtex
	case "katex":
		mathrenderingengine = convert.Katex
	case "gladtex":
		mathrenderingengine = convert.Gladtex
	default:
		mathrenderingengine = convert.Mathml
	}
	convert.SetMathRenderingOptions(mathrenderingengine, config.MathRenderingURL)
}

func PreflightConfiguration(config cfgh.Config) {
	log.Info("preflight applying configuration")
	err := os.Setenv("PANDOC_COMMAND", config.PandocCommand)
	err = os.Setenv("LATEX_COMMAND", config.LatexCommand)
	if !config.UseGoTex {
		log.Infof("disabling Gotex")
		convert.DisableGoTex()
	} else {
		convert.EnableGoTex()
	}
	if err != nil {
		log.Errorf("Error setting Environment Variables: %s", err)
	}
	convert.SetTimeout(config.Timeout)
	setMathRenderingEngine(config)
	go func() {
		_, err := zero.Register(config)
		if err != nil {
			log.Errorf("Error registering zero plugin: %s", err)
		}
	}()
}

func PreflightConfigCheck(config cfgh.Config) {
	log.Info("preflight checking configuration")
	if config.TrustedProxy == "0.0.0.0" {
		log.Warn("trusted proxy is set to 0.0.0.0. This is considered UNSAFE. Please set a trusted proxy either by setting ´TRUSTED_PROXY´ or by using the ´-trust-proxy´ command line option.")
	}
	if config.Debug {
		log.Warn("debug mode is enabled, you might want to disable it in production environments. To disable debug mode unset ´DEBUG´ (and/or check if you use the `-debug` command line option).")
	}
	if len(config.PandocCommand) == 0 {
		log.Fatal("pandoc command is not defined. Please set it with ´PANDOC_COMMAND´ environment variable or by using the ´-pandoc-command´ command line option.")
	}
}

func PreflightChecks() {
	log.Info("preflight checks")
	determineVersionOfPackageAndSetEnv(os.Getenv("PANDOC_COMMAND"), "PANDOC_VERSION")
}

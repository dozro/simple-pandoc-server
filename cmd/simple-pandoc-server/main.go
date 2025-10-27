package main

import (
	"simple-pandoc-server/internal/pkg/checks"
	cfgh "simple-pandoc-server/internal/pkg/confighandling"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	checks.PreflightPackageSearch()
	cfg := cfgh.ReadConfigFromEnv()
	if cfg.Debug {
		log.SetLevel(log.DebugLevel)
	} else {
		gin.SetMode(gin.ReleaseMode)
		log.SetLevel(log.InfoLevel)
	}
	checks.PreflightConfigCheck(cfg)
	checks.PreflightConfiguration(cfg)
	checks.PreflightChecks()
	log.Info("starting server")
	log.Debugf("Config: %+v", cfg)
	startServer(cfg)
}

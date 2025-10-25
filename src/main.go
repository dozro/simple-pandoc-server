package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	PreflightPackageSearch()
	cfg := readConfigFromEnv()
	if cfg.Debug {
		log.SetLevel(log.DebugLevel)
	} else {
		gin.SetMode(gin.ReleaseMode)
		log.SetLevel(log.InfoLevel)
	}
	PreflightConfigCheck(cfg)
	PreflightConfiguration(cfg)
	PreflightChecks()
	log.Info("starting server")
	log.Debugf("Config: %+v", cfg)
	startServer(cfg)
}

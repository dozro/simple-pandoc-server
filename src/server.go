package main

import (
	"simple-pandoc-server/convert"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func startServer(cfg Config) {
	log.Debugf("starting web server on %s", cfg.ListenOnIP)
	router := gin.Default()
	err := router.SetTrustedProxies([]string{cfg.trustedProxy})
	if err != nil {
		log.Fatal(err)
	}
	// latex
	router.POST("/parse/latex/toPdf/raw", convert.ParseLatexRawToPDF)
	router.POST("/parse/latex/toHtml/raw", convert.ParseLatexRawToHtml)
	router.POST("/parse/latex/toHtml/plain", convert.ParseLatexPlainToHtml)
	router.POST("/parse/latex/toPdf/plain", convert.ParseLatexPlainToPdf)
	// typst
	router.POST("/parse/typst/toHtml/raw", convert.ParseTypstRawToHtml)
	// health
	router.GET("/healt", getHealth)
	err = router.Run(cfg.ListenOnIP)
	if err != nil {
		log.Fatal(err)
	}
}

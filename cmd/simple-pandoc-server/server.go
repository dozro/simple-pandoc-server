package main

import (
	cfgh "simple-pandoc-server/internal/pkg/confighandling"
	"simple-pandoc-server/internal/pkg/server"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func startServer(cfg cfgh.Config) {
	log.Debugf("starting web server on %s", cfg.ListenOnIP)
	router := gin.Default()
	err := router.SetTrustedProxies([]string{cfg.TrustedProxy})
	if err != nil {
		log.Fatal(err)
	}
	// latex
	router.POST("/parse/latex/toPdf/raw", server.ParseLatexRawToPDF)
	router.POST("/parse/latex/toHtml/raw", server.ParseLatexRawToHTML)
	// docx
	router.POST("/parse/docx/toPdf/raw", server.ParseDocxRawToPDF)
	router.POST("/parse/docx/toHtml/raw", server.ParseDocxRawToHTML)
	// odt
	router.POST("/parse/odt/toPdf/raw", server.ParseOdtRawToPdf)
	router.POST("/parse/odt/toHtml/raw", server.ParseOdtRawToHTML)
	// typst
	router.POST("/parse/typst/toHtml/raw", server.ParseTypstRawToHtml)
	router.POST("/parse/typst/toPdf/raw", server.ParseTypstRawToPdf)
	// health
	router.GET("/health", getHealth)
	err = router.Run(cfg.ListenOnIP)
	if err != nil {
		log.Fatal(err)
	}
}

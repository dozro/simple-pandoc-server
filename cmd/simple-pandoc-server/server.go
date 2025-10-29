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
	//router.POST("/parse/latex/toHtml/plain", server.ParseLatexRawToHTML)
	//router.POST("/parse/latex/toPdf/plain", convert.ParseLatexPlainToPdf)
	// docx
	router.POST("/parse/docx/toPdf/raw", server.ParseDocxRawToPDF)
	router.POST("/parse/docx/toHtml/raw", server.ParseDocxRawToHTML)
	// typst
	router.POST("/parse/typst/toHtml/raw", server.ParseTypstRawToHtml)
	// health
	router.GET("/health", getHealth)
	err = router.Run(cfg.ListenOnIP)
	if err != nil {
		log.Fatal(err)
	}
}

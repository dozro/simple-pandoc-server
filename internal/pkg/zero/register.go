package zero

import (
	cfgh "simple-pandoc-server/internal/pkg/confighandling"
	"strconv"
	"strings"

	"github.com/grandcat/zeroconf"
	log "github.com/sirupsen/logrus"
)

func Register(cfg cfgh.Config) {
	var add []string
	add = strings.Split(cfg.ListenOnIP, ":")
	if len(add) < 2 {
		log.Fatalf("ListenOnIP value '%s' is malformed, expected format 'host:port'", cfg.ListenOnIP)
	}
	port, err := strconv.Atoi(add[1])
	if err != nil {
		log.Fatalf("Invalid port in ListenOnIP: %v", err)
	}
	_, err = zeroconf.Register("pandoc", "_http._tcp", add[0], port, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
}

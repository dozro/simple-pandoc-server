package zero

import (
	cfgh "simple-pandoc-server/internal/pkg/confighandling"
	"strconv"
	"strings"

	"github.com/grandcat/zeroconf"
	log "github.com/sirupsen/logrus"
)

func Register(cfg cfgh.Config) (*zeroconf.Server, error) {
	var add []string
	add = strings.Split(cfg.ListenOnIP, ":")
	if len(add) < 2 {
		log.Fatalf("ListenOnIP value '%s' is malformed, expected format 'host:port'", cfg.ListenOnIP)
	}
	port, err := strconv.Atoi(add[1])
	if err != nil {
		log.Errorf("Invalid port in ListenOnIP: %v", err)
		return nil, err
	}
	server, err := zeroconf.Register("pandoc", "_http._tcp", add[0], port, nil, nil)
	if err != nil {
		return nil, err
	}
	return server, nil
}

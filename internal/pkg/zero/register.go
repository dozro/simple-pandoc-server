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
	port, _ := strconv.Atoi(add[1])
	_, err := zeroconf.Register("pandoc", "_http._tcp", add[0], port, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
}

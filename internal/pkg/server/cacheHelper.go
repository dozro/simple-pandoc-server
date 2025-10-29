package server

import (
	"crypto/md5"
	"fmt"
	"simple-pandoc-server/internal/pkg/cache"

	log "github.com/sirupsen/logrus"
)

func hashOfData(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

func toCache(data []byte, out []byte) {
	go func() {
		dataHash := hashOfData(data)
		_ = cache.AddToCache(dataHash, out)
	}()
}
func fromCache(data []byte) ([]byte, error) {
	dataHash := hashOfData(data)
	out, exists := cache.GetFromCache(dataHash)
	if !exists {
		log.Debugf("cache miss for %s", dataHash)
		return nil, fmt.Errorf("data not found in cache for hash: %s", dataHash)
	}
	log.Debugf("cache hit for %s", dataHash)
	bytesOut, ok := out.([]byte)
	if !ok {
		return nil, fmt.Errorf("cached value for hash %s is not []byte (got %T)", dataHash, out)
	}
	log.Debugf("cache content for %s: %s", dataHash, bytesOut)

	return bytesOut, nil
}

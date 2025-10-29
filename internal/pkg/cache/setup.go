package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

func SetupCache(expireAfter time.Duration, cleanupInterval time.Duration) {
	c = cache.New(expireAfter, cleanupInterval)
}

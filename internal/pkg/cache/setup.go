package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var c *cache.Cache
var expiry time.Duration

func SetupCache(expireAfter time.Duration, cleanupInterval time.Duration) {
	expiry = expireAfter
	c = cache.New(expireAfter, cleanupInterval)
}

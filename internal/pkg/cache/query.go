package cache

import "fmt"

func GetFromCache(cacheKey string) (interface{}, bool) {
	if c == nil {
		return nil, false
	}
	return c.Get(cacheKey)
}
func AddToCache(cacheKey string, data []byte) error {
	if c == nil {
		return fmt.Errorf("cache not initialized")
	}
	return c.Add(cacheKey, data, expiry)
}

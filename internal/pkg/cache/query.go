package cache

func GetFromCache(cacheKey string) (interface{}, bool) {
	return c.Get(cacheKey)
}
func AddToCache(cacheKey string, data []byte) error {
	return c.Add(cacheKey, data, expiry)
}

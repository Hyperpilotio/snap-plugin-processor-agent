package agent

// Cache in-memory store to keep metrics
type Cache struct {
	Data map[string]CacheType
}

// CacheType define what kind of attribute of metrics should keep
type CacheType struct {
	Pre interface{}
}

// NewCache return an instance of Cache
func NewCache() *Cache {
	return &Cache{
		Data: make(map[string]CacheType),
	}
}

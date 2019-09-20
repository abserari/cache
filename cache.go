package acache

import "sync"

var (
	cache = make(map[string]*CacheTable)
	mutex sync.RWMutex
)

func Cache(table string) *CacheTable {
	mutex.RLock()
	t, ok := cache[table]
	mutex.RUnlock()

	if !ok {
		mutex.Lock()
		// double check
		t, ok = cache[table]
		t = &CacheTable{
			name:  table,
			items: make(map[interface{}]*CacheItem),
		}
		cache[table] = t
	}
	mutex.Unlock()

	return t
}

package acache

import (
	"sync"
	"time"
)

// CacheItem -
type CacheItem struct {
	sync.RWMutex

	key   interface{}
	value interface{}
	// live lifeSpan time
	lifeSpan time.Duration

	createdAt   time.Time
	accessedAt  time.Time
	accessCount int64

	expireTrigger []func(key interface{})
}

func NewCacheItem(key interface{}, lifeSpan time.Duration, value interface{}) *CacheItem {
	t := time.Now()
	return &CacheItem{
		key:      key,
		value:    value,
		lifeSpan: lifeSpan,

		createdAt:   t,
		accessedAt:  t,
		accessCount: 0,

		expireTrigger: nil,
	}
}

func (item *CacheItem) KeepAlive() {
	item.Lock()
	defer item.Unlock()
	item.accessedAt = time.Now()
	item.accessCount++
}

// Getter

func (item *CacheItem) LifeSpan() time.Duration {
	return item.lifeSpan
}

func (item *CacheItem) AccessCount() int64 {
	return item.accessCount
}

func (item *CacheItem) AccessedAt() time.Time {
	return item.accessedAt
}

func (item *CacheItem) CreatedAt() time.Time {
	return item.createdAt
}

func (item *CacheItem) Key() interface{} {
	return item.key
}

func (item *CacheItem) Value() interface{} {
	return item.value
}

// Trigger func

func (item *CacheItem) SetExpireTrigger(f func(interface{})) {
	if len(item.expireTrigger) > 0 {
		item.RemoveExpireTrigger()
	}

	item.Lock()
	defer item.Unlock()
	item.expireTrigger = append(item.expireTrigger, f)
}

func (item *CacheItem) AddExpireTrigger(f func(interface{})) {
	item.Lock()
	defer item.Unlock()
	item.expireTrigger = append(item.expireTrigger, f)
}

func (item *CacheItem) RemoveExpireTrigger() {
	item.Lock()
	defer item.Unlock()
	item.expireTrigger = nil
}

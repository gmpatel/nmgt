package datasource

import (
	"sync"
	"time"
)

type item struct {
	value      interface{}
	lastAccess int64
}

// FreqCache holds freq accessed key for faster response
type FreqCache struct {
	data map[string]*item
	dc   DataSource
	sync.Mutex
}

// NewFreqCache returns
func NewFreqCache(dc DataSource, cacheSeconds int) (fc *FreqCache) {
	mycache := new(FreqCache) // Can do this way as well &FreqCache{data: make(map[string]*item)}
	mycache.data = make(map[string]*item)
	mycache.dc = dc

	go func() {
		for now := range time.Tick(time.Second) {
			mycache.Lock()
			for key, val := range mycache.data {
				if now.Unix()-val.lastAccess > int64(cacheSeconds) {
					delete(mycache.data, key)

				}
			}
			mycache.Unlock()
		}
	}()

	return mycache
}

// Value ...
func (ds *FreqCache) Value(key string) (interface{}, error) {
	ds.Lock()
	vald, keyExists := ds.data[key]
	if keyExists {
		ds.Unlock()
		return vald.value, nil
	}
	ds.Unlock()

	valc, err := ds.dc.Value(key)
	if err != nil {
		return nil, err
	}

	ds.Lock()
	ds.data[key] = &item{
		value:      valc,
		lastAccess: time.Now().Unix(),
	}
	ds.Unlock()
	return valc, nil
}

// Store ...
func (ds *FreqCache) Store(key string, val interface{}) error {
	ds.Lock()
	ds.data[key] = &item{
		value:      val,
		lastAccess: time.Now().Unix(),
	}
	ds.Unlock()

	// Syncing value back to db
	go func() {
		ds.dc.Store(key, val)
	}()

	return nil
}

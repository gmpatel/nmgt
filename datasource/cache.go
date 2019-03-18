package datasource

import (
	"sync"
	"time"
)

// DistributedCache ...
type DistributedCache struct {
	data map[string]interface{}
	db   DataSource
	sync.Mutex
}

// NewDistributedCache ...
func NewDistributedCache(db DataSource) *DistributedCache {
	dc := new(DistributedCache)
	dc.data = make(map[string]interface{})
	dc.db = db
	return dc
}

// Value ...
func (dc *DistributedCache) Value(key string) (interface{}, error) {
	// simulate 100ms roundtrip to the distributed cache
	time.Sleep(100 * time.Millisecond)

	/*
		So, to bring the response time below "100ms" I have implemented another layer "FreqCache" from the "DataSource" interface
		where I'am storing key/value of the frequently accessed keys.
	*/

	dc.Lock()

	value, keyExists := dc.data[key]
	if keyExists {
		dc.Unlock()
		return value, nil
	}

	value, err := dc.db.Value(key)
	if err != nil {
		dc.Unlock()
		return nil, err
	}

	dc.data[key] = value
	dc.Unlock()

	return value, nil
}

// Store ...
func (dc *DistributedCache) Store(key string, value interface{}) error {
	// simulate 100ms roundtrip to the distributed cache
	time.Sleep(100 * time.Millisecond)

	dc.Lock()
	dc.data[key] = value
	dc.Unlock()

	// Syncing value back to db
	go func() {
		dc.db.Store(key, value)
	}()

	return nil
}

// GetData is added only for the debug/dev purpose to check the cache storage map "state" directly if require as "data" is private and not exported
// GetData ideally should be removed when not needed
func (dc *DistributedCache) GetData() map[string]interface{} {
	return dc.data
}

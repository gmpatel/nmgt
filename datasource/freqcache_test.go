package datasource

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFreqCacheBehavior(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	initDbRecords := 10
	queryKey := "key1"
	expectedValue := "value1"

	// Setup
	db := NewDatabase()
	db.initData(initDbRecords)
	dc := NewDistributedCache(db)
	ds := NewFreqCache(dc, 1)

	// Perform
	start := time.Now()
	value, err := ds.Value(queryKey)
	elapsed := time.Since(start)

	start1 := time.Now()
	value1, err1 := ds.Value(queryKey)
	elapsed1 := time.Since(start1)

	time.Sleep(3 * time.Second)

	start2 := time.Now()
	value2, err2 := ds.Value(queryKey)
	elapsed2 := time.Since(start2)

	// Assertions
	assert.Equal(t, len(db.data), initDbRecords)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, expectedValue, value)
	assert.GreaterOrEqual(t, int64(elapsed), int64(600*time.Millisecond))

	assert.Nil(t, err1)
	assert.NotNil(t, value1)
	assert.Equal(t, expectedValue, value1)
	assert.LessOrEqual(t, int64(elapsed1), int64(100*time.Millisecond))

	assert.Nil(t, err2)
	assert.NotNil(t, value2)
	assert.Equal(t, expectedValue, value2)
	assert.GreaterOrEqual(t, int64(elapsed2), int64(100*time.Millisecond))
}

func TestFreqCacheStore(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	initDbRecords := 10
	queryKey := "key10"
	expectedValue := "value10"

	// Setup
	db := NewDatabase()
	db.initData(initDbRecords)
	dc := NewDistributedCache(db)
	ds := NewFreqCache(dc, 1)

	// Perform
	err := ds.Store(queryKey, expectedValue)
	time.Sleep(1 * time.Second) // this will allow to update the DistributedCache & Database with newly inserted key in FreqCache (there can be better way of doing this)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, len(db.data), initDbRecords+1)
	assert.Equal(t, len(dc.data), 1)
	assert.Equal(t, len(ds.data), 1)
}

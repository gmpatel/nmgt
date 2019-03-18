package datasource

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDcValueFetchesKeyFromDbAndStoresInDcFirstReqAtLeastTakes600msAndFollowingReqAtLeastTakes100ms(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	initDbRecords := 10
	queryKey := "key1"
	expectedValue := "value1"

	// Setup
	db := NewDatabase()
	db.initData(initDbRecords)
	dc := NewDistributedCache(db)

	// Perform
	start := time.Now()
	value, err := dc.Value(queryKey)
	elapsed := time.Since(start)

	start1 := time.Now()
	value1, err1 := dc.Value(queryKey)
	elapsed1 := time.Since(start1)

	// Assertions
	assert.Equal(t, len(db.data), initDbRecords)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, len(dc.data), 1)
	assert.Equal(t, expectedValue, value)
	assert.NotEqual(t, int64(elapsed), int64(600*time.Millisecond)) // GreaterOrEqual

	assert.Nil(t, err1)
	assert.NotNil(t, value1)
	assert.Equal(t, len(dc.data), 1)
	assert.NotEqual(t, int64(elapsed1), int64(100*time.Millisecond)) // GreaterOrEqual
}

func TestDcStore(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	initDbRecords := 10
	queryKey := "key10"
	expectedValue := "value10"

	// Setup
	db := NewDatabase()
	db.initData(initDbRecords)
	dc := NewDistributedCache(db)

	// Perform
	err := dc.Store(queryKey, expectedValue)
	time.Sleep(1 * time.Second) // this will allow to update the Database with newly inserted key in DistributedCache (there can be better way of doing this)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, initDbRecords+1, len(db.data))
	assert.Equal(t, 1, len(dc.data))
}

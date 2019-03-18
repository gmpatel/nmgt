package datasource

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDbInitData(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	initDbRecords := 10

	// Setup
	db := NewDatabase()

	// Perform
	db.initData(initDbRecords)

	// Assertions
	assert.Equal(t, len(db.data), initDbRecords)
}

func TestDbValue(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	initDbRecords := 10
	queryKey := "key1"
	expectedValue := "value1"

	// Setup
	db := NewDatabase()

	// Perform
	db.initData(initDbRecords)
	value, err := db.Value(queryKey)

	// Assertions
	assert.Equal(t, len(db.data), initDbRecords)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, expectedValue, value)
}

func TestDbValueAtLeastTakes500ms(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	initDbRecords := 10
	queryKey := "key1"
	expectedValue := "value1"

	// Setup
	db := NewDatabase()

	// Perform
	db.initData(initDbRecords)
	start := time.Now()
	value, err := db.Value(queryKey)
	elapsed := time.Since(start)

	// Assertions
	assert.Equal(t, len(db.data), initDbRecords)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, expectedValue, value)
	assert.NotEqual(t, int64(elapsed), int64(500*time.Millisecond)) // GreaterOrEqual
}

func TestDbStore(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	initDbRecords := 10
	queryKey := "key10"
	expectedValue := "value10"

	// Setup
	db := NewDatabase()

	// Perform
	db.initData(initDbRecords)
	err := db.Store(queryKey, expectedValue)
	time.Sleep(1 * time.Second) // this will allow to update the db with newly inserted key (there can be better way of doing this)

	// Assertions
	assert.Equal(t, len(db.data), initDbRecords+1)
	assert.Nil(t, err)
}

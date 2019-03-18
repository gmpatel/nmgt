package datasource

import (
	"fmt"
	"sync"
	"time"
)

// Database struct is an object contains logic to store and retrive data from Databse
type Database struct {
	data map[string]interface{}
	sync.Mutex
}

// NewDatabase constructor to be used for actual production code
func NewDatabase() *Database {
	db := new(Database)
	db.data = make(map[string]interface{})
	return db
}

// NewDatabaseWithInitData constructor to be used for testing and to init db with some data initially
func NewDatabaseWithInitData(count int) *Database {
	db := new(Database)
	db.data = make(map[string]interface{})
	db.initData(count)
	return db
}

// Value returns the value for the provided key from the db store
func (db *Database) Value(key string) (interface{}, error) {
	// simulate 500ms roundtrip to the distributed cache
	time.Sleep(500 * time.Millisecond)

	value, keyExists := db.data[key]
	if keyExists {
		return value, nil
	}

	return nil, fmt.Errorf("key['%s'] does not exists", key)
}

// Store stores the value for the provided key into the db store
func (db *Database) Store(key string, value interface{}) error {
	// simulate 500ms roundtrip to the distributed cache
	time.Sleep(500 * time.Millisecond)

	db.Lock()
	db.data[key] = value
	db.Unlock()

	return nil	
}

// GetData is added only for the debug/dev purpose to check the database storage map "state" directly if require as "data" is private and not exported
// GetData ideally should be removed when not needed
func (db *Database) GetData() map[string]interface{} {
	return db.data
}

func (db *Database) initData(count int) {
	for i := 0; i < count; i++ {
		key := fmt.Sprintf("key%d", i)
		val := fmt.Sprintf("value%d", i)

		/*
			I can use "db.Store(key, val)" method to initialise data
			to the database but that will take time with the simulation
			of the 500m. But in real life the data will already be there
			in db, so, I can either insert staright to the map as well.
		*/

		// db.Store(key, val)

		/*
			So to simulate data already exist in the db I am instead
			directly storing values to the map!
		*/

		db.data[key] = val
	}
}

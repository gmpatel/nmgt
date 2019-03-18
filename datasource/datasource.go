package datasource

// DataSource interface
type DataSource interface {
	Value(key string) (interface{}, error)

	/*
		Adding "Store(..).." method to interface as I am injecting "Database"
		to the "DistributedCache" as a "DataSource interface" instead of
		concrete "Database" object so that we can easily replace the end "Database"
		to something else...

		I am not sure I am supposed to update the "interface" or not for this exercise, but, doing it
		for the good cause so that we can easily chage DataSource with diff implementations...
	*/
	Store(key string, value interface{}) error
}

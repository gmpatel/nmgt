package main

import (
	"github.com/gmpatel/nmgt/benchmarking"
	"github.com/gmpatel/nmgt/datasource"
	"github.com/gmpatel/nmgt/processors"
)

func main() {
	var processor processors.Processor

	// Performace Processor configuration, which can be moved to the config file as well
	threads := 10           // will trigger 10 Go routines for testing concurrency
	requestsPerThread := 50 // each Go routine will send 50 iterative requests
	fastCacheSeconds := 5   // fast Cache will hold data for 5 seconds from last accessed point for fast retrival of freq accessed data
	initDataSetCount := 10  // will initialize DB with 10 Key/Value paors

	/*
		Creating Performace Processor with given configuration. Performace Processor has been implemented on
		"Processor" interface so it can be easlily replaced by another processor or components.
	*/

	ab := benchmarking.NewAverageBenchmarking()
	db := datasource.NewDatabaseWithInitData(initDataSetCount)
	dc := datasource.NewDistributedCache(db)
	ds := datasource.NewFreqCache(dc, fastCacheSeconds)

	processor = processors.NewPerformanceProcessor(
		threads,
		requestsPerThread,
		ds,
		ab,
	)
	processor.Process()
}

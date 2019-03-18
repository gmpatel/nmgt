package processors

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gmpatel/nmgt/benchmarking"
	"github.com/gmpatel/nmgt/datasource"
)

// PerformanceProcessor object processes concorunt processes to benchmark datasource performace
type PerformanceProcessor struct {
	ds                datasource.DataSource
	ab                benchmarking.Benchmarking
	threads           int
	requestsPerThread int
}

// NewPerformanceProcessor returns the processor object of the program for main
func NewPerformanceProcessor(threads int, requestsPerThread int, ds datasource.DataSource, ab benchmarking.Benchmarking) *PerformanceProcessor {
	return &PerformanceProcessor{
		threads:           threads,
		requestsPerThread: requestsPerThread,
		ds:                ds,
		ab:                ab,
	}
}

// Process is the
func (pc *PerformanceProcessor) Process() {
	var wg sync.WaitGroup

	for i := 0; i < pc.threads; i++ {
		wg.Add(1)
		go func(gorn int) {
			defer wg.Done()
			for counter := 0; counter < pc.requestsPerThread; counter++ {
				rnd := pc.random(0, 9)
				key := fmt.Sprintf("key%d", rnd)
				pc.getKeyValueAndLogTimeTrack(key, gorn, counter)
			}
		}(i)
	}

	wg.Wait()
	pc.logPerformanceBenchmark()
}

func (pc *PerformanceProcessor) getKeyValueAndLogTimeTrack(key string, gorn int, counter int) (interface{}, error) {
	start := time.Now()

	value, err := pc.getKeyValue(key)
	if err != nil {
		pc.logTimeTrack(start, key, err, gorn, counter)
		return nil, err
	}

	pc.logTimeTrack(start, key, value, gorn, counter)
	return value, nil
}

func (pc *PerformanceProcessor) getKeyValue(key string) (interface{}, error) {
	value, err := pc.ds.Value(key)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (pc *PerformanceProcessor) random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func (pc *PerformanceProcessor) logTimeTrack(start time.Time, key string, value interface{}, gorn int, counter int) {
	elapsed := time.Since(start)
	flag := ""

	pc.ab.Add(elapsed)

	message := fmt.Sprintf(">> go-func[%02d][%03d] Request '%s', response '%s', time:", gorn+1, counter+1, key, value)
	if elapsed < 100*time.Millisecond {
		flag = " (< 100ms)"
	}
	log.Printf("%s %s%s", message, elapsed, flag)
}

func (pc *PerformanceProcessor) logPerformanceBenchmark() {
	_, calls, bm, unit := pc.ab.Benchmark("ms")
	log.Printf(">> Performance Benchmark (Average): %d Calls, %f %s/call (BM)", calls, bm, unit)
}

package benchmarking

import (
	"sync"
	"time"
)

// AverageBenchmarking is an implementation of "Benchmarking" interface to calculate average of the requests response
type AverageBenchmarking struct {
	total int64
	count int
	sync.Mutex
}

// NewAverageBenchmarking constroctor returns the AverageBenchmarking object
func NewAverageBenchmarking() *AverageBenchmarking {
	ab := new(AverageBenchmarking)
	return ab
}

// Add keep adding durations took
func (ab *AverageBenchmarking) Add(duration time.Duration) {
	ab.Lock()
	defer ab.Unlock()

	ab.total += duration.Nanoseconds()
	ab.count++
}

// Benchmark returns calculated benchmark in this case average of all logged calls
func (ab *AverageBenchmarking) Benchmark(code string) (int64, int, float64, string) {
	switch code {
	case "ns":
		return ab.total, ab.count, float64(ab.total) / float64(ab.count), "ns"
	case "μs":
		return ab.total, ab.count, float64(ab.total) / (float64(ab.count) * 1000), "μs"
	case "ms":
		return ab.total, ab.count, float64(ab.total) / (float64(ab.count) * 1000000), "ms"
	default:
		return ab.total, ab.count, float64(ab.total) / (float64(ab.count) * 1000000000), "s"
	}
}

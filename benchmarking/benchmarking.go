package benchmarking

import "time"

// Benchmarking is an interface to provide different benchmarking mechanism
type Benchmarking interface {
	Add(duration time.Duration)
	Benchmark(code string) (int64, int, float64, string)
}

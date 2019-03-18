package benchmarking

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAverageBenchmarkingAddDurationCorrectly(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	requestsCount := 10
	eachRequestDuration := 100 * time.Millisecond

	// Setup
	averageBenchmarking := NewAverageBenchmarking()

	// Perform
	for i := 0; i < requestsCount; i++ {
		averageBenchmarking.Add(eachRequestDuration)
	}

	// Assertions
	assert.Equal(t, averageBenchmarking.total, int64(requestsCount*int(eachRequestDuration)))
}

func TestAverageBenchmarkingCalculatesBenchmarkCorrectly(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	requestsCount := 10
	eachRequestDuration := 100 * time.Millisecond
	unit := "ms"

	// Setup
	averageBenchmarking := NewAverageBenchmarking()

	// Perform
	for i := 0; i < requestsCount; i++ {
		averageBenchmarking.Add(eachRequestDuration)
	}
	duration, count, avg, u := averageBenchmarking.Benchmark(unit)

	// Assertions
	assert.Equal(t, requestsCount, count)
	assert.Equal(t, unit, u)
	assert.Equal(t, eachRequestDuration.Nanoseconds()*int64(requestsCount), duration)
	assert.Equal(t, int(eachRequestDuration/time.Millisecond), int(avg))
}

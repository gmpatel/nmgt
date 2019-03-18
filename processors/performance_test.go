package processors

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDataSource is an autogenerated mock type for the storager type
type MockDataSource struct {
	mock.Mock
}

// Value returns the value for the key provided
func (ds *MockDataSource) Value(key string) (interface{}, error) {
	args := ds.Called(key)
	return args.Get(0).(interface{}), args.Error(1)
}

// Store stores the value provided for the key
func (ds *MockDataSource) Store(key string, value interface{}) error {
	args := ds.Called(key, value)
	return args.Error(0)
}

// MockBenchmarking is an autogenerated mock type for the benchmarking
type MockBenchmarking struct {
	mock.Mock
}

// Add adds the response time of the request to the total
func (ds *MockBenchmarking) Add(duration time.Duration) {
	ds.Called(duration)
	return
}

// Benchmark returns the benchmarking info for the all requests
func (ds *MockBenchmarking) Benchmark(code string) (int64, int, float64, string) {
	args := ds.Called(code)
	return args.Get(0).(int64), args.Int(1), args.Get(2).(float64), args.String(3)
}

// GetDefaultTestPerformanceProcessor returns "processors.PerformanceProcessor" with default configruation
func GetDefaultTestPerformanceProcessor(threads int, requestsPerThread int) (*PerformanceProcessor, *MockDataSource, *MockBenchmarking) {
	// Setup
	ds := &MockDataSource{}
	ds.On("Value", mock.Anything).Return("value", nil)
	ds.On("Store", mock.Anything, mock.Anything).Return(nil)

	ab := &MockBenchmarking{}
	ab.On("Add", mock.Anything).Return()
	ab.On("Benchmark", mock.Anything).Return(int64(500*time.Millisecond), 500, 1.00, "ms")

	// Perform
	performanceProcessor := NewPerformanceProcessor(
		threads,
		requestsPerThread,
		ds,
		ab,
	)
	return performanceProcessor, ds, ab
}

func TestRandomInPerformanceProcessor(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	threads := 10
	requestsPerThread := 50

	// Setup
	performanceProcessor, _, _ := GetDefaultTestPerformanceProcessor(
		threads,
		requestsPerThread,
	)

	// Perform
	rand := performanceProcessor.random(1, 5)

	// Assertions
	assert.GreaterOrEqual(t, rand, 1)
	assert.LessOrEqual(t, rand, 10)
}

func TestGetKeyValueInPerformanceProcessorTriggersDataSourceValue(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	threads := 10
	requestsPerThread := 50

	// Setup
	performanceProcessor, ds, _ := GetDefaultTestPerformanceProcessor(
		threads,
		requestsPerThread,
	)

	// Perform
	val, err := performanceProcessor.getKeyValue("key1")

	// Assertions
	assert.NotNil(t, val)
	assert.Nil(t, err)
	assert.Equal(t, val, "value")

	ds.AssertNumberOfCalls(t, "Value", 1)
}

func TestGetKeyValueWithLogInPerformanceProcessorTriggersDataSourceValue(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	threads := 10
	requestsPerThread := 50

	// Setup
	performanceProcessor, ds, _ := GetDefaultTestPerformanceProcessor(
		threads,
		requestsPerThread,
	)

	val, err := performanceProcessor.getKeyValueAndLogTimeTrack("key1", 1, 1)

	// Assertions
	assert.NotNil(t, val)
	assert.Nil(t, err)
	assert.Equal(t, val, "value")

	ds.AssertNumberOfCalls(t, "Value", 1)
}

func TestLogTimeTrackInPerformanceProcessorTriggersBenchmarkingAdd(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	threads := 10
	requestsPerThread := 50

	// Setup
	performanceProcessor, _, ab := GetDefaultTestPerformanceProcessor(
		threads,
		requestsPerThread,
	)

	// Perform
	performanceProcessor.logTimeTrack(time.Now(), "key1", "value", 1, 1)

	// Assertions
	ab.AssertNumberOfCalls(t, "Add", 1)
}

func TestLogPerformanceBenchmarkInPerformanceProcessorTriggersBenchmarkingBenchmark(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	threads := 10
	requestsPerThread := 50

	// Setup
	performanceProcessor, _, ab := GetDefaultTestPerformanceProcessor(
		threads,
		requestsPerThread,
	)

	// Perform
	performanceProcessor.logPerformanceBenchmark()

	// Assertions
	ab.AssertNumberOfCalls(t, "Benchmark", 1)
}

func TestPerformanceProcessorInterfaceProcessorProcessMethod(t *testing.T) {
	assert.New(t)

	// Configuration & Expectations
	threads := 10
	requestsPerThread := 50
	expectedDataAccessCalls := threads * requestsPerThread
	expectedBenchmarkingAddCalls := threads * requestsPerThread
	expectedBenchmarkingBenchmarkCalls := 1

	// Setup
	pp, ds, ab := GetDefaultTestPerformanceProcessor(
		threads,
		requestsPerThread,
	)

	// Perform
	pp.Process()

	// Assertion
	ds.AssertNumberOfCalls(t, "Value", expectedDataAccessCalls)
	ab.AssertNumberOfCalls(t, "Add", expectedBenchmarkingAddCalls)
	ab.AssertNumberOfCalls(t, "Benchmark", expectedBenchmarkingBenchmarkCalls)
}

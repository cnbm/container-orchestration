package generic

// Result represents the results of a benchmark run
type Result struct {
}

// BenchmarkRun represents a single run of a container orchestration benchmark
type BenchmarkRun interface {
	Setup() error
	Execute() (Result, error)
	Teardown() error
}

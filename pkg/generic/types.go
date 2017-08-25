package generic

// BenchmarkTarget represents a container orchestration system that is the target of a benchmark run
type BenchmarkTarget string

// BenchmarkRunner represents a single run of a container orchestration benchmark against a concrete target
type BenchmarkRunner interface {
	Setup() error
	Execute() (Result, error)
	Teardown() error
}

// Result represents the results of a benchmark run
type Result struct {
}
package generic

// BenchmarkRunType represents a benchmark run type as of https://github.com/cnbm/container-orchestration#benchmark-run-types
type BenchmarkRunType string

// BenchmarkTarget represents a container orchestration system that is the target of a benchmark run
type BenchmarkTarget string

// BenchmarkRunner represents a single run of a container orchestration benchmark against a concrete target
type BenchmarkRunner interface {
	Setup() error
	Execute() (BenchmarkResult, error)
	Teardown() error
}

// BenchmarkResult represents the results of a benchmark run
type BenchmarkResult struct {
	Output string
}

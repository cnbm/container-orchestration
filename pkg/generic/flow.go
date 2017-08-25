package generic

import (
	"time"
)

// Run is the generic benchmark flow
func Run(b BenchmarkRunner) (time.Duration, error) {
	err := b.Setup()
	if err != nil {
		return 0, err
	}
	startTime := time.Now()
	result, err := b.Execute()
	if err != nil {
		return 0, err
	}
	_ = result
	elapsed := time.Since(startTime)
	err = b.Teardown()
	if err != nil {
		return 0, err
	}
	return elapsed, nil
}

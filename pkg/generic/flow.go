package generic

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

// Run is the generic benchmark flow
func Run(b BenchmarkRunner) (BenchmarkResult, time.Duration, error) {
	setupErr := b.Setup()
	if setupErr != nil {
		// Setup error, it does not make sense to run benchmark.
		return BenchmarkResult{}, 0, setupErr
	}

	startTime := time.Now()
	result, executeErr := b.Execute()
	_ = result
	elapsed := time.Since(startTime)

	teardownErr := b.Teardown()
	if teardownErr != nil {
		log.Error("Error while tearing down benchmark: %v", teardownErr)
	}

	return result, elapsed, executeErr
}

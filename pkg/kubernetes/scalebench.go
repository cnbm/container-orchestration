package kubernetes

import (
	log "github.com/Sirupsen/logrus"
	"github.com/cnbm/container-orchestration/pkg/generic"
)

// Scalebench represents the Kubernetes specific benchmark run for the scaling benchmark
type Scalebench struct {
	Config map[string]string
}

// Setup prepares and inits the Kubernetes environment for the scaling benchmark
func (bench Scalebench) Setup() error {
	log.Info("Setting up Kubernetes scaling benchmark")
	return nil
}

// Execute executes the scaling benchmark against a Kubernetes cluster
func (bench Scalebench) Execute() (generic.Result, error) {
	log.Info("Executing Kubernetes scaling benchmark")
	r := generic.Result{}
	// bench.Config["apiserver"]
	return r, nil
}

// Teardown tears down and cleans up the Kubernetes environment after the scaling benchmark has executed
func (bench Scalebench) Teardown() error {
	log.Info("Tearing down Kubernetes scaling benchmark")
	return nil
}

package kubernetes

import (
	"fmt"
	"os/exec"

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
func (bench Scalebench) Execute() (generic.BenchmarkResult, error) {
	log.Info("Executing Kubernetes scaling benchmark")
	r := generic.BenchmarkResult{}
	output, err := exec.Command("kubectl", "version").CombinedOutput()
	if err != nil {
		return r, fmt.Errorf("Failed to shell out:%s", err)
	}
	r.Output = string(output)
	return r, nil
}

// Teardown tears down and cleans up the Kubernetes environment after the scaling benchmark has executed
func (bench Scalebench) Teardown() error {
	log.Info("Tearing down Kubernetes scaling benchmark")
	return nil
}

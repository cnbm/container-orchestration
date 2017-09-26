package kubernetes

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/cnbm/container-orchestration/pkg/generic"
)

// ServiceDiscovery represents the Kubernetes specific benchmark run for the servicediscovery benchmark
type ServiceDiscovery struct {
	Config map[string]string
}

// Setup prepares and inits the Kubernetes environment for the servicediscovery benchmark
func (bench ServiceDiscovery) Setup() error {
	log.Info("Setting up Kubernetes servicediscovery benchmark")
	cs, err := getclient(bench.Config["kubeconfig"])
	if err != nil {
		return err
	}
	sisepod := gensisep(bench.Config["ns"])
	_, err = cs.CoreV1().Pods(bench.Config["ns"]).Create(sisepod)
	if err != nil {
		return fmt.Errorf("Can't create pod 'sise': %s", err)
	}
	return nil
}

// Execute executes the servicediscovery benchmark against a Kubernetes cluster
func (bench ServiceDiscovery) Execute() (generic.BenchmarkResult, error) {
	log.Info("Executing Kubernetes servicediscovery benchmark")
	r := generic.BenchmarkResult{}
	cs, err := getclient(bench.Config["kubeconfig"])
	if err != nil {
		return r, err
	}
	_ = cs
	sise := gensises(bench.Config["ns"])
	s, err := cs.Services(bench.Config["ns"]).Create(sise)
	if err != nil {
		return r, fmt.Errorf("Can't create service: %s", err)
	}
	_ = s
	// wait until 'sise.cnbm.svc/info' returns 200:
	// svcdone(cs, ns string, s)
	r.Output = "benchmark succeeded"
	return r, nil
}

// Teardown tears down and cleans up the Kubernetes environment after the servicediscovery benchmark has executed
func (bench ServiceDiscovery) Teardown() error {
	log.Info("Tearing down Kubernetes servicediscovery benchmark")
	return nil
}

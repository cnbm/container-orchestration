package kubernetes

import (
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
	// s, err := cs.Services(v1.NamespaceDefault).Create(*v1.Service)
	// if err != nil {
	// 	return r, fmt.Errorf("Can't create service: %s", err)
	// }
	// r.Output = fmt.Sprintf("%v", s)
	return r, nil
}

// Teardown tears down and cleans up the Kubernetes environment after the servicediscovery benchmark has executed
func (bench ServiceDiscovery) Teardown() error {
	log.Info("Tearing down Kubernetes servicediscovery benchmark")
	return nil
}

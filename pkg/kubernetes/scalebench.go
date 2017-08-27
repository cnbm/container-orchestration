package kubernetes

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/cnbm/container-orchestration/pkg/generic"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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
	config, err := clientcmd.BuildConfigFromFlags("", bench.Config["kubeconfig"])
	if err != nil {
		return r, fmt.Errorf("Failed to build config for Kubernetes: %s", err)
	}
	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		return r, fmt.Errorf("Failed to create client for Kubernetes: %s", err)
	}
	d, err := cs.ExtensionsV1beta1().Deployments("").List(metav1.ListOptions{})
	if err != nil {
		return r, fmt.Errorf("Can't get deployments: %s", err)
	}
	r.Output = "No deployments found"
	if len(d.Items) > 0 {
		r.Output = fmt.Sprintf("Found %d deployments", len(d.Items))
	}
	return r, nil
}

// Teardown tears down and cleans up the Kubernetes environment after the scaling benchmark has executed
func (bench Scalebench) Teardown() error {
	log.Info("Tearing down Kubernetes scaling benchmark")
	return nil
}

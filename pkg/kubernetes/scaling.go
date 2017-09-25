package kubernetes

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/cnbm/container-orchestration/pkg/generic"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Scaling represents the Kubernetes specific benchmark run for the scaling benchmark
type Scaling struct {
	Config map[string]string
}

// Setup prepares and inits the Kubernetes environment for the scaling benchmark
func (bench Scaling) Setup() error {
	log.Info("Setting up Kubernetes scaling benchmark")
	return nil
}

// Execute executes the scaling benchmark against a Kubernetes cluster
func (bench Scaling) Execute() (generic.BenchmarkResult, error) {
	log.Info("Executing Kubernetes scaling benchmark")
	r := generic.BenchmarkResult{}
	cs, err := getclient(bench.Config["kubeconfig"])
	if err != nil {
		return r, err
	}
	busybox := genbusyboxd(
		bench.Config["numpods"],
		bench.Config["cpu"],
		bench.Config["mem"],
	)
	d, err := cs.AppsV1beta1().Deployments(bench.Config["ns"]).Create(busybox)
	if err != nil {
		return r, fmt.Errorf("Can't create deployment 'cnbm-co-scaling': %s", err)
	}
	deploydone(cs, bench.Config["ns"], d, func(i string) { log.Infof("Deployment done: %s", i) })
	r.Output = fmt.Sprintf("%v", busybox)
	return r, nil
}

// Teardown tears down and cleans up the Kubernetes environment after the scaling benchmark has executed
func (bench Scaling) Teardown() error {
	log.Info("Tearing down Kubernetes scaling benchmark")
	cs, err := getclient(bench.Config["kubeconfig"])
	if err != nil {
		return err
	}
	err = cs.AppsV1beta1().Deployments(bench.Config["ns"]).Delete("cnbm-co-scaling", &metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("Can't delete deployment 'cnbm-co-scaling': %s", err)
	}
	return nil
}

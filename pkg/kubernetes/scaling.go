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
	deploydone(cs, bench.Config["ns"], d)
	r.Output = "benchmark succeeded"
	return r, nil
}

// Teardown tears down and cleans up the Kubernetes environment after the scaling benchmark has executed
func (bench Scaling) Teardown() error {
	log.Info("Tearing down Kubernetes scaling benchmark")
	cs, err := getclient(bench.Config["kubeconfig"])
	if err != nil {
		return err
	}
	d, err := cs.AppsV1beta1().Deployments(bench.Config["ns"]).Get("cnbm-co-scaling", metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("Can't get deployment 'cnbm-co-scaling': %s", err)
	}
	d.Spec.Replicas = int32Ptr("0")
	_, err = cs.AppsV1beta1().Deployments(bench.Config["ns"]).Update(d)
	if err != nil {
		return fmt.Errorf("Can't scale down deployment 'cnbm-co-scaling': %s", err)
	}
	prop := metav1.DeletePropagationForeground
	err = cs.AppsV1beta1().Deployments(bench.Config["ns"]).Delete("cnbm-co-scaling", &metav1.DeleteOptions{PropagationPolicy: &prop})
	if err != nil {
		return fmt.Errorf("Can't delete deployment 'cnbm-co-scaling': %s", err)
	}
	return nil
}

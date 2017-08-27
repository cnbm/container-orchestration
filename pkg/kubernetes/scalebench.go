package kubernetes

import (
	"fmt"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/cnbm/container-orchestration/pkg/generic"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/apps/v1beta1"
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
	busybox := genbusyboxd(
		bench.Config["numpods"],
		bench.Config["cpu"],
		bench.Config["mem"],
	)
	_, err = cs.AppsV1beta1().Deployments(v1.NamespaceDefault).Create(busybox)
	if err != nil {
		return r, fmt.Errorf("Can't deploy busybox: %s", err)
	}
	r.Output = fmt.Sprintf("%v", busybox)
	return r, nil
}

// Teardown tears down and cleans up the Kubernetes environment after the scaling benchmark has executed
func (bench Scalebench) Teardown() error {
	log.Info("Tearing down Kubernetes scaling benchmark")
	return nil
}

// genbusyboxd creates a deployment with numpods pods and each
// with resource constraints (limit==request) which must be at least
// 1 millicore for cpuusagesec (effectively: "0.001") and
// 4MB for meminbytes (effectively: "4000000")
func genbusyboxd(numpods, cpuusagesec, meminbytes string) *v1beta1.Deployment {
	return &v1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "bench",
		},
		Spec: v1beta1.DeploymentSpec{
			Replicas: int32Ptr(numpods),
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "cnbm-co",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "busybox",
							Image: "busybox",
							Command: []string{
								"sleep",
								"10000",
							},
							Resources: tolimits(meminbytes, cpuusagesec),
							Ports: []v1.ContainerPort{
								{
									Name:          "http",
									Protocol:      v1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
}

func tolimits(cpuusagesec, meminbytes string) v1.ResourceRequirements {
	cpuval, _ := resource.ParseQuantity(cpuusagesec)
	memval, _ := resource.ParseQuantity(meminbytes)
	newlim := v1.ResourceList{
		v1.ResourceCPU:    cpuval,
		v1.ResourceMemory: memval,
	}
	return v1.ResourceRequirements{
		Limits:   newlim,
		Requests: newlim,
	}
}

func int32Ptr(i string) *int32 {
	v, _ := strconv.ParseInt(i, 10, 32)
	v32 := int32(v)
	return &v32
}

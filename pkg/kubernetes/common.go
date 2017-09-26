package kubernetes

import (
	"fmt"
	"strconv"
	"time"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/apps/v1beta1"
	"k8s.io/client-go/tools/clientcmd"
)

func getclient(confloc string) (*kubernetes.Clientset, error) {
	cs := &kubernetes.Clientset{}
	config, err := clientcmd.BuildConfigFromFlags("", confloc)
	if err != nil {
		return cs, fmt.Errorf("Failed to build config for Kubernetes: %s", err)
	}
	cs, err = kubernetes.NewForConfig(config)
	if err != nil {
		return cs, fmt.Errorf("Failed to create client for Kubernetes: %s", err)
	}
	return cs, nil
}

func deploydone(cs *kubernetes.Clientset, ns string, dorig *v1beta1.Deployment) {
	want := dorig.Spec.Replicas
	for {
		time.Sleep(100 * time.Millisecond)
		d, err := cs.AppsV1beta1().Deployments(ns).Get(dorig.GetName(), metav1.GetOptions{})
		if err != nil {
			continue
		}
		have := d.Status.AvailableReplicas
		if *want == have {
			return
		}
	}
}

// genbusyboxd creates a deployment in namespace ns with numpods pods and each
// with resource constraints (limit==request) which must be at least
// 1 millicore for cpuusagesec (effectively: "0.001") and
// 4MB for meminbytes (effectively: "4000000")
func genbusyboxd(ns, numpods, cpuusagesec, meminbytes string) *v1beta1.Deployment {
	return &v1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cnbm-co-scaling",
			Namespace: ns,
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
							Resources: tolimits(cpuusagesec, meminbytes),
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

// tolimits sets the resources requirements and limits
// to the respective parameters cpuusagesec and meminbytes
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

// gensisep creates a pod using mhausenblas/simpleservice:0.5.0 image
func gensisep(ns string) *v1.Pod {
	p := &v1.Probe{
		Handler: v1.Handler{
			HTTPGet: &v1.HTTPGetAction{
				Path: "/health",
				Port: intstr.FromInt(9876),
			},
		},
	}

	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "sise",
			Namespace: ns,
			Labels: map[string]string{
				"app": "cnbm-co",
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "sise",
					Image: "mhausenblas/simpleservice:0.5.0",
					Ports: []v1.ContainerPort{
						{ContainerPort: 9876},
					},
					LivenessProbe:  p,
					ReadinessProbe: p,
				},
			},
		},
	}
}

// gensises creates a service using app=cnbm-co selector
func gensises(ns string) *v1.Service {
	return &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "sise",
			Namespace: ns,
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{Port: 80, TargetPort: intstr.FromInt(9876)},
			},
			Selector: map[string]string{
				"app": "cnbm-co",
			},
		},
	}
}

func int32Ptr(i string) *int32 {
	v, _ := strconv.ParseInt(i, 10, 32)
	v32 := int32(v)
	return &v32
}

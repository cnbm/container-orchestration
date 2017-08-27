# Container Orchestration Benchmark

[![Go Report Card](https://goreportcard.com/badge/github.com/cnbm/container-orchestration)](https://goreportcard.com/report/github.com/cnbm/container-orchestration)
[![godoc](https://godoc.org/github.com/cnbm/container-orchestration/pkg?status.svg)](https://godoc.org/github.com/cnbm/container-orchestration/pkg)

The purpose of the container orchestration benchmark (`cnbm-co` for short) is to provide a vendor-neutral, extendable benchmark for container orchestration systems. The current focus is on stateless workloads and we're implementing it for the following container orchestration systems (targets):

- DC/OS
- Kubernetes

If you want to contribute, simply fork this repo, add your implementation in `pkg/` and send in a [PR](https://github.com/cnbm/container-orchestration/pulls).

Contents:

- [Using](#using)
  - [Launching](#launching)
  - [Availability matrix](#availability-matrix)
- [Developing](#developing)
- [Design](design.md)

## Using

### Launching

In general:

```
$ ./cnbm-co launch -h
Launches the CNBM container orchestration benchmark

Usage:
  cnbm-co launch [flags]

Flags:
  -h, --help             help for launch
  -p, --params string    Comma separated key-value pair list of target-specific configuration parameters. For example: k1=v1,k2=v2
  -r, --runtype string   The benchmark run type. Allowed values: [scaling distribution apicalls servicediscovery recovery]
  -t, --target string    The target container orchestration system to benchmark. Allowed values: [dcos kubernetes]

Global Flags:
      --config string   config file (default is $HOME/.cnbm.yaml)
```

#### DC/OS

```
$ ./cnbm-co launch --runtype scaling --target dcos -p dcosurl=http://example.com,dcosacstoken=123
INFO[0000] Setting up DC/OS scaling benchmark
INFO[0000] Executing DC/OS scaling benchmark
INFO[0000] Deploying a new application
INFO[0000] RESULT:
 Target: DC/OS
 Output: {}
 Elapsed time: 0s
```

#### Kubernetes

To benchmark a Kubernetes cluster, use `--target kubernetes`, for example, to launch the `scaling` run type, do:

```
$ ./cnbm-co launch --runtype scaling --target kubernetes --params kubeconfig=/Users/mhausenblas/.kube/config,numpods=1,cpu=0.1,mem=67108864
INFO[0000] Setting up Kubernetes scaling benchmark
INFO[0000] Executing Kubernetes scaling benchmark
INFO[0000] Tearing down Kubernetes scaling benchmark
INFO[0000] RESULT:
 Target: Kubernetes
 Output: {&Deployment{ObjectMeta:k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{Name:bench,GenerateName:,Namespace:,SelfLink:,UID:,ResourceVersion:,Generation:0,CreationTimestamp:0001-01-01 00:00:00 +0000 UTC,DeletionTimestamp:<nil>,DeletionGracePeriodSeconds:nil,Labels:map[string]string{},Annotations:map[string]string{},OwnerReferences:[],Finalizers:[],ClusterName:,Initializers:nil,},Spec:DeploymentSpec{Replicas:*10,Selector:nil,Template:k8s_io_kubernetes_pkg_api_v1.PodTemplateSpec{ObjectMeta:k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{Name:,GenerateName:,Namespace:,SelfLink:,UID:,ResourceVersion:,Generation:0,CreationTimestamp:0001-01-01 00:00:00 +0000 UTC,DeletionTimestamp:<nil>,DeletionGracePeriodSeconds:nil,Labels:map[string]string{app: cnbm-co,},Annotations:map[string]string{},OwnerReferences:[],Finalizers:[],ClusterName:,Initializers:nil,},Spec:PodSpec{Volumes:[],Containers:[{busybox busybox [sleep 10000] []  [{http 0 80 TCP }] [] [] {map[cpu:{{0 0} {<nil>}  } memory:{{0 0} {<nil>}  }] map[cpu:{{0 0} {<nil>}  } memory:{{0 0} {<nil>}  }]} [] nil nil nil    nil false false false}],RestartPolicy:,TerminationGracePeriodSeconds:nil,ActiveDeadlineSeconds:nil,DNSPolicy:,NodeSelector:map[string]string{},ServiceAccountName:,DeprecatedServiceAccount:,NodeName:,HostNetwork:false,HostPID:false,HostIPC:false,SecurityContext:nil,ImagePullSecrets:[],Hostname:,Subdomain:,Affinity:nil,SchedulerName:,InitContainers:[],AutomountServiceAccountToken:nil,Tolerations:[],HostAliases:[],},},Strategy:DeploymentStrategy{Type:,RollingUpdate:nil,},MinReadySeconds:0,RevisionHistoryLimit:nil,Paused:false,RollbackTo:nil,ProgressDeadlineSeconds:nil,},Status:DeploymentStatus{ObservedGeneration:0,Replicas:0,UpdatedReplicas:0,AvailableReplicas:0,UnavailableReplicas:0,Conditions:[],ReadyReplicas:0,CollisionCount:nil,},}}
 Elapsed time: 17.206911ms
```

Note the following params that are required, depending on the `--runtype`:

- `scaling`
  - `kubeconfig` … the Kubernetes config to use, for example `/Users/mhausenblas/.kube/config`
  - `numpods` … the number of pods to launch, for example, `10`
  - `cpu` … the CPU seconds (at least `0.01`) per pod
  - `mem` … the memory (at least `4000000`) per pod
- TBD


### Availability matrix

The following matrix shows the availability of [benchmark run types](design.md#benchmark-run-types) per [target](design.md#targets):

| benchmark run type   | DC/OS    | Kubernetes |
| --------------------:| -------- | ---------- |
| `scaling`            | Y        | Y          |
| `distribution`       | N        | N          |
| `distribution`       | N        | N          |
| `apicalls`           | N        | N          |
| `servicediscovery`   | N        | N          |

## Developing

### Building

```
$ make
Building the CNBM CO CLI
go build -ldflags "-X github.com/cnbm/container-orchestration/cli/cmd.releaseVersion=0.1.0" -o ./cnbm-co cli/main.go
```

### Vendoring

We are using Go [dep](https://github.com/golang/dep) for dependency management. If you don't have `dep` installed yet, do `go get -u github.com/golang/dep/cmd/dep` now and then:

```
$ dep ensure
```

### Testing

For unit tests we use the `go test` command, for example:

```
$ go test -v -short -run Test* .
```

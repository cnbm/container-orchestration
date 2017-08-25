# Container Orchestration Benchmark

[![Go Report Card](https://goreportcard.com/badge/github.com/cnbm/container-orchestration)](https://goreportcard.com/report/github.com/cnbm/container-orchestration)
[![godoc](https://godoc.org/github.com/cnbm/container-orchestration/pkg?status.svg)](https://godoc.org/github.com/cnbm/container-orchestration/pkg)

The purpose of the container orchestration benchmark (`cnbm-co` for short) is to provide a vendor-neutral, extendable benchmark for container orchestration systems. The current focus is on stateless workloads and we're implementing it for the following container orchestration systems (targets):

- DC/OS
- Kubernetes

If you want to contribute, simply fork this repo, add your implementation in `pkg/` and send in a [PR](https://github.com/cnbm/container-orchestration/pulls).

Contents:

- [Using a benchmark](#using-a-benchmark)
- [Developing](#developing)
- [Benchmark design](#benchmark-design)
- [Related Work](#related-work)

## Using a benchmark

### Launching

```
$ ./cnbm-co launch -h
Launches the CNBM container orchestration benchmark

Usage:
  cnbm-co launch [flags]

Flags:
  -h, --help            help for launch
  -p, --params string   Comma separated key-value pair list of target-specific configuration parameters. For example: k1=v1,k2=v2
  -t, --target string   The target container orchestration system to benchmark. Allowed values: [dcos k8s]

Global Flags:
      --config string   config file (default is $HOME/.cnbm.yaml)

$ ./cnbm-co launch -t dcos -p dcosurl=http://example.com,dcosacstoken=123
INFO[0000] Setting up DC/OS scale test
INFO[0000] Executing DC/OS scale test
INFO[0000] Deploying a new application
INFO[0000] Elapsed time for the scaling benchmark for DC/OS: 1s
```

### Availability matrix

The following matrix shows the availability of [benchmark run types](#benchmark-run-types) per [target](#targets):

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

## Benchmark design

### Targets

The benchmark is executed as follows:

- User provisions the cluster and provides a running cluster to the benchmark.
- Benchmark itself runs in the the cluster, triggered by the local `cnbm-co` command.
- Results are dumped to stdout as CSV/JSON, locally.

Supported targets:

- [DC/OS 1.9.2](https://dcos.io/releases/1.9.2/)
- [Kubernetes 1.7.4](https://github.com/kubernetes/kubernetes/releases/tag/v1.7.4)

### Benchmark run types

#### `scaling`

The following sequence:

1. Start `N` containers in `seconds` potentially with different runtimes (Docker, UCR, CRI-O).
1. Stop `N` containers in `seconds`.

#### `distribution`

Launches `N` containers and measures the distribution over nodes in `map: nodeid -> set of containers`.

#### `apicalls`

Measures API calls from within cluster in `seconds`:

- list containers
- list pods
- list services/endpoints

#### `servicediscovery`

Measure service discovery in `seconds`:

- Start a service and measure how long it takes until it can be discovered from different nodes.
- How long does a query/look-up take (while scaling services)?

#### `recovery`

Recovery performance in case of re-scheduling a pod/ (container) in  `seconds`.

### Dimensions

For each benchmark run, the following dimensions should be considered (where applicable):

- Number nodes, that is, worker nodes that are hosting containers
- Number of containers
- Container runtime type (Docker, UCR, CRI-O)
- Failure rate (per container, nodes, network)

## Related Work

- openshift/svt [cluster-loader](https://github.com/openshift/svt/tree/master/openshift_scalability)
- [C4-bench](https://github.com/allingeek/c4-bench)
- [Go-based framework for running benchmarks against Docker, containerd, and runc engine layers](https://github.com/estesp/bucketbench)
- [1000 nodes and beyond: updates to Kubernetes performance and scalability in 1.2](http://blog.kubernetes.io/2016/03/1000-nodes-and-beyond-updates-to-Kubernetes-performance-and-scalability-in-12.html)
- [OpenShift v3 Scaling, Performance and Capacity Planning](https://access.redhat.com/articles/2191731)
- [Deploying 1000 nodes of OpenShift on the CNCF Cluster (Part 1)](https://www.cncf.io/blog/2016/08/23/deploying-1000-nodes-of-openshift-on-the-cncf-cluster-part-1)
- [Exploring Performance of etcd, Zookeeper and Consul Consistent Key-value Datastores](https://coreos.com/blog/performance-of-etcd.html)

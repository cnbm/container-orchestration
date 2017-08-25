# Container Orchestration Benchmark

The purpose of the container orchestration benchmark (`cnbm-co` for short) is to provide a vendor-neutral, extendable benchmark for container orchestration systems. The current focus is on stateless workloads and we're implementing it for the following container orchestration systems:

- DC/OS
- Kubernetes

If you want to contribute, simply fork this repo, add your implementation in `pkg/` and send in a PR.

- [Benchmark Design](#benchmark-design-and-setup)
- [Using a benchmark](#using-a-benchmark)
- [Developing](#developing)
- [Related Work](#related-work)

## Benchmark Design

### Targets

- Start <n> container [seconds]
    - Docker (prefetched)
    - UCR
    - CRI-O
- Stop <n> container [seconds]
- Container distribution over nodes [Map (nodeid -> container)]
- API calls from within cluster [seconds]
    - List containers
- Service Discovery [seconds]
    - Start 1 service, how long until it can be discovered from different nodes
    - How long does query/look-up take (while scaling services)?
- recovery performance (in case of re-scheduling)

### Dimensions

For each run, the following dimensions can be considered:

- number nodes (hosting containers)
- number of containers
- container runtimes (Docker, rkt, UCR, etc.)
- failure rate
  - container
  - nodes
  - network

### Flow
  - User provides a running cluster
  - Benchmark itself runs in cluster (Docker run, Marathon spec, K8S spec), triggered from local environment
  - Results are dumped in CSV/JSON       

## Using a benchmark

```
$ ./cnbm-co launch -t dcos -p dcosurl=http://example.com,dcosacstoken=123
INFO[0000] Setting up DC/OS scale test
INFO[0000] Executing DC/OS scale test
INFO[0000] Deploying a new application
INFO[0000] Elapsed time for the scaling benchmark for DC/OS: 1s
```

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

## Related Work

- openshift/svt [cluster-loader](https://github.com/openshift/svt/tree/master/openshift_scalability)
- [C4-bench](https://github.com/allingeek/c4-bench)
- [Go-based framework for running benchmarks against Docker, containerd, and runc engine layers](https://github.com/estesp/bucketbench)
- [1000 nodes and beyond: updates to Kubernetes performance and scalability in 1.2](http://blog.kubernetes.io/2016/03/1000-nodes-and-beyond-updates-to-Kubernetes-performance-and-scalability-in-12.html)
- [OpenShift v3 Scaling, Performance and Capacity Planning](https://access.redhat.com/articles/2191731)
- [Deploying 1000 nodes of OpenShift on the CNCF Cluster (Part 1)](https://www.cncf.io/blog/2016/08/23/deploying-1000-nodes-of-openshift-on-the-cncf-cluster-part-1)
- [Exploring Performance of etcd, Zookeeper and Consul Consistent Key-value Datastores](https://coreos.com/blog/performance-of-etcd.html)

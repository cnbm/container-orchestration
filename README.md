# Container Orchestration Benchmark

The purpose of the Container Orchestration Benchmark (cnbm-cob for short) is to provide a vendor-neutral, extendable benchmark for container orchestration systems. The current focus is on stateless workloads.

## Setup

### Targets

- Start <n> container [seconds]
    - Docker Container (prefetched),
    - UCR/CRYO
- Stop <n> Container [seconds]
- Container Distribution over nodes [Map (nodeid -> container)]
- API calls from within cluster [seconds]
    - List Container 
- Service Discovery [seconds]
    - Start 1 service, how long until it can be   discovered from different nodes
    - How long does query take (while scaling services)
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
  - Benchmark itself runs in cluster (docker run, marathon json), triggered from local environment
  - Results are dumped in CSV/JSON       

## Dependencies

- [Cobra](https://github.com/spf13/cobra)
- add here ...

## Related Work

- openshift/svt [cluster-loader](https://github.com/openshift/svt/tree/master/openshift_scalability)
- [C4-bench](https://github.com/allingeek/c4-bench)
- [Go-based framework for running benchmarks against Docker, containerd, and runc engine layers](https://github.com/estesp/bucketbench)
- [1000 nodes and beyond: updates to Kubernetes performance and scalability in 1.2](http://blog.kubernetes.io/2016/03/1000-nodes-and-beyond-updates-to-Kubernetes-performance-and-scalability-in-12.html)
- [OpenShift v3 Scaling, Performance and Capacity Planning](https://access.redhat.com/articles/2191731)
- [Deploying 1000 nodes of OpenShift on the CNCF Cluster (Part 1)](https://www.cncf.io/blog/2016/08/23/deploying-1000-nodes-of-openshift-on-the-cncf-cluster-part-1)
- [Exploring Performance of etcd, Zookeeper and Consul Consistent Key-value Datastores](https://coreos.com/blog/performance-of-etcd.html)

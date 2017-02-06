# Container Orchestration Benchmark

The purpose of the Container Orchestration Benchmark (cnbm-cob for short) is to provide a vendor-neutral, extendable benchmark for container orchestration systems. The current focus is on stateless workloads.

## Setup

### Targets

- start-up time of containers
- tear-down time of containers
- containers distribution over nodes
- external API responsiveness
- service discovery performance
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

## Related Work

- [C4-bench](https://github.com/allingeek/c4-bench)
- [1000 nodes and beyond: updates to Kubernetes performance and scalability in 1.2](http://blog.kubernetes.io/2016/03/1000-nodes-and-beyond-updates-to-Kubernetes-performance-and-scalability-in-12.html)
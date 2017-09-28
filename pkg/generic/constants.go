package generic

const (
	// DefaultOutputDir is the default output directory for results
	DefaultOutputDir = "/tmp/cnbm/results"
	
	
	// TargetDCOS represents DC/OS as the target container orchestration system
	TargetDCOS BenchmarkTarget = "dcos"
	// TargetK8S represents Kubernetes as the target container orchestration system
	TargetK8S BenchmarkTarget = "kubernetes"
	
	
	// RunScaling represents the scaling benchmark run type
	// Precondition: 
	//	* empty cluster
	//	* busybox:1 docker image prefetched on all nodes
	// Benchmark: 
	//	* scale-up 3 container on empty cluster
	// Container Specification
	//	* docker image: busybox:1 (prefetched), 
	//      * container runtime: free to choose, potential parameter 
	//      * cpu: 0.5
	//      * memory: 300MB
	// Postcondition: all container running (i.e., not necessarily ready/passing health check)
	// Result: start up time in seconds
	RunScaling BenchmarkRunType = "scaling"
	
	// RunDistribution represents the distribution benchmark run type
	RunDistribution BenchmarkRunType = "distribution"
	// RunAPICalls represents the apicalls benchmark run type
	RunAPICalls BenchmarkRunType = "apicalls"
	// RunSD represents the servicediscovery benchmark run type
	RunSD BenchmarkRunType = "servicediscovery"
	// RunRecovery represents the recovery benchmark run type
	RunRecovery BenchmarkRunType = "recovery"
)

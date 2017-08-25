package generic

const (
	// DefaultOutputDir is the default output directory for results
	DefaultOutputDir = "/tmp/cnbm/results"
	// TargetDCOS represents DC/OS as the target container orchestration system
	TargetDCOS BenchmarkTarget = "dcos"
	// TargetK8S represents Kubernetes as the target container orchestration system
	TargetK8S BenchmarkTarget = "kubernetes"
	// RunScaling represents the scaling benchmark run type
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

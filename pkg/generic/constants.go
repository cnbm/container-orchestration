package generic

const (
	// DefaultOutputDir is the default output directory for results
	DefaultOutputDir = "/tmp/cnbm/results"
	// TargetDCOS represents DC/OS as the target container orchestration system
	TargetDCOS BenchmarkTarget = "dcos"
	// TargetK8S represents Kubernetes as the target container orchestration system
	TargetK8S BenchmarkTarget = "k8s"
)

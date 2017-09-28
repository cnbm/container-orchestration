package cmd

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cnbm/container-orchestration/pkg/dcos"
	"github.com/cnbm/container-orchestration/pkg/generic"
	"github.com/cnbm/container-orchestration/pkg/kubernetes"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launches the CNBM container orchestration benchmark",
	Run: func(cmd *cobra.Command, args []string) {
		// process and validate flags:
		r := generic.BenchmarkRunType(cmd.Flag("runtype").Value.String())
		if r == "" {
			errornexit("No benchmark run type provided")
		}
		t := generic.BenchmarkTarget(cmd.Flag("target").Value.String())
		if t == "" {
			errornexit("No target provided")
		}
		paramlistraw := cmd.Flag("params").Value.String()
		if paramlistraw == "" {
			errornexit("No configuration parameters for target provided")
		}
		paramlist := strings.Split(paramlistraw, ",")
		configmap := map[string]string{}
		for _, kvraw := range paramlist {
			kv := strings.Trim(kvraw, " ")
			k := strings.Split(kv, "=")[0]
			v := strings.Split(kv, "=")[1]
			configmap[k] = v
		}
		// determine target and init the benchmark run type accordingly:
		var s generic.BenchmarkRunner
		var targetname string
		switch t {
		case generic.TargetDCOS:
			switch r {
			case generic.RunScaling:
				s = dcos.Scaling{Config: configmap}
			case generic.RunSD:
				s = dcos.ServiceDiscovery{Config: configmap}
			case generic.RunAPICalls:
				s = dcos.ApiCall{Config: configmap}
			case generic.RunDistribution:
				s = dcos.Distribution{Config: configmap}
			case generic.RunRecovery:
				s = dcos.Recovery{Config: configmap}
			default:
				errornexit("Benchmark run type unknown")
			}
			targetname = "DC/OS"
			//TODO: check if I got all the necessary DC/OS config parameters such as URL ("dcosurl") and ACS token ("dcosacstoken")
		case generic.TargetK8S:
			switch r {
			case generic.RunScaling:
				s = kubernetes.Scaling{Config: configmap}
			case generic.RunSD:
				s = kubernetes.ServiceDiscovery{Config: configmap}
			default:
				errornexit("Benchmark run type unknown")
			}
			targetname = "Kubernetes"
		default:
			errornexit("Target unknown")
		}
		// run the parameterized benchmark:
		result, elapsed, err := generic.Run(s)
		if err != nil {
			errornexit(fmt.Sprintf("Wasn't able to run benchmark for %s: %s", targetname, err))
		}
		log.Infof("RESULT:\n Target: %s\n Output: %s\n Elapsed time: %v", targetname, result, elapsed)
	},
}

func init() {
	RootCmd.AddCommand(launchCmd)
	runtypes := []generic.BenchmarkRunType{
		generic.RunScaling,
		generic.RunDistribution,
		generic.RunAPICalls,
		generic.RunSD,
		generic.RunRecovery,
	}
	launchCmd.Flags().StringP("runtype", "r", "", fmt.Sprintf("The benchmark run type. Allowed values: %v", runtypes))
	targets := []generic.BenchmarkTarget{
		generic.TargetDCOS,
		generic.TargetK8S,
	}
	launchCmd.Flags().StringP("target", "t", "", fmt.Sprintf("The target container orchestration system to benchmark. Allowed values: %v", targets))
	launchCmd.Flags().StringP("params", "p", "", "Comma separated key-value pair list of target-specific configuration parameters. For example: k1=v1,k2=v2")
}

func errornexit(msg string) {
	log.Fatalf(fmt.Sprintf("%s. Exiting …", msg))
}

package cmd

import (
	"fmt"
	"os"
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
		t := generic.BenchmarkTarget(cmd.Flag("target").Value.String())
		if t == "" {
			log.Errorf("No target provided, exiting …")
			os.Exit(1)
		}
		paramlistraw := cmd.Flag("params").Value.String()
		if paramlistraw == "" {
			log.Error("No configuration parameters for target provided, exiting …")
			os.Exit(1)
		}
		paramlist := strings.Split(paramlistraw, ",")
		configmap := map[string]string{}
		for _, kvraw := range paramlist {
			kv := strings.Trim(kvraw, " ")
			k := strings.Split(kv, "=")[0]
			v := strings.Split(kv, "=")[1]
			configmap[k] = v
		}
		// determine target and init accordingly the benchmark(s):
		var s generic.BenchmarkRunner
		var targetname string
		switch t {
		case generic.TargetDCOS:
			s = dcos.Scalebench{Config: configmap}
			targetname = "DC/OS"
			//TODO: check if I got all the necessary DC/OS config parameters such as URL ("dcosurl") and ACS token ("dcosacstoken")
		case generic.TargetK8S:
			s = kubernetes.Scalebench{Config: configmap}
			targetname = "Kubernetes"
			//TODO: check if I got all the necessary K8S config parameters
		default:
			log.Error("Target unknown, try something else")
		}
		// run the parameterized benchmark:
		result, elapsed, err := generic.Run(s)
		if err != nil {
			log.Errorf("Wasn't able to run benchmark for %s: %s", targetname, err)
		}
		log.Infof("RESULT:\n Target: %s\n Output: %s\n Elapsed time: %v", targetname, result, elapsed)
	},
}

func init() {
	RootCmd.AddCommand(launchCmd)
	targets := []generic.BenchmarkTarget{generic.TargetDCOS, generic.TargetK8S}
	launchCmd.Flags().StringP("target", "t", "", fmt.Sprintf("The target container orchestration system to benchmark. Allowed values: %v", targets))
	launchCmd.Flags().StringP("params", "p", "", "Comma separated key-value pair list of target-specific configuration parameters. For example: k1=v1,k2=v2")
}

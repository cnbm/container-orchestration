package cmd

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cnbm/container-orchestration/pkg/dcos"
	"github.com/cnbm/container-orchestration/pkg/generic"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launches the CNBM container orchestration benchmark",
	Run: func(cmd *cobra.Command, args []string) {
		t := generic.BenchmarkTarget(cmd.Flag("target").Value.String())
		configlist := strings.Split(cmd.Flag("config").Value.String(), ",")
		configmap := map[string]string{}
		for _, kvraw := range configlist {
			kv := strings.Trim(kvraw, " ")
			k := strings.Split(kv, "=")[0]
			v := strings.Split(kv, "=")[1]
			configmap[k] = v
		}
		switch t {
		case generic.TargetDCOS:
			launchDCOS(configmap)
		case generic.TargetK8S:
			log.Info("Not implemented yet")
		default:
			log.Error("Target unknown, try something else")
		}
	},
}

func init() {
	RootCmd.AddCommand(launchCmd)
	targets := []generic.BenchmarkTarget{generic.TargetDCOS, generic.TargetK8S}
	launchCmd.PersistentFlags().StringP("target", "t", "", fmt.Sprintf("The target container orchestration system to benchmark. Allowed values: %v", targets))
	_ = launchCmd.MarkFlagRequired("target")
	launchCmd.PersistentFlags().StringP("config", "c", "", "A comma separated key-value pair list of target-specific configuration parameters, for example the cluster API and a token: api=http://api.example.com,token=12345")
	_ = launchCmd.MarkFlagRequired("config")
}

func launchDCOS(cm map[string]string) {
	s := dcos.Scalebench{Config: cm}
	//TODO: check if I got all the necessary DC/OS config parameters: URL ("dcosurl") and ACS token ("dcosacstoken")
	elapsed, err := generic.Run(s)
	if err != nil {
		log.Errorf("There was a problem carrying out the scaling benchmark for DC/OS: %s", err)
	}
	log.Info("Elapsed time for the scaling benchmark for DC/OS: %v", elapsed)
}

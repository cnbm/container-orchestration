package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"

	dcos "github.com/cnbm/container-orchestration/pkg/dcos"
	"github.com/cnbm/container-orchestration/pkg/generic"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launches the CNBM container orchestration benchmark",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Executing the CNBM container orchestration benchmark")
		s := dcos.Scalebench{
			DCOSURL:      "https://joerg-22r-elasticl-mv9wyg0lclf4-218935880.us-west-2.elb.amazonaws.com",
			DCOSACSToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1aWQiOiJib290c3RyYXB1c2VyIiwiZXhwIjoxNTAyMDg5NjAyfQ.gMxn7mw7om5tdsBJS6qtDzD6_mUR1ySzNndB0JS2ZiIUGOE6lYDFB2K22uWFi_hqKRg3RPHUdJI47esvY0DlWH20veLSDJEA9vRzg9qPLcKXzrCy_zwF_q1fw_uwkEIdVrmvttHmNEWiW4V1bbDajx9lWDFiKhz7d7p5BHaYvP1ycDhVTjDTBLyBAIC4CAdgFoh1MFocXfNk-SC4yXp68H4v13bTL7jhwjHpgeRWK_c2NH8J53vJUJdOuXqnTqNMKdbZ0D03kx5AlaNdMWpTiAMteschn9ZsdlaeihKLoqvPGPQR-emuOsua0h0njWobqAUhMVcsoere0eJF1_l5dg",
		}
		elapsed, err := generic.Run(s)
		if err != nil {
			log.Errorf("There was a problem carrying out the benchmark: %s", err)
		}
		log.Info("Elapsed time: %v", elapsed)
	},
}

func init() {
	RootCmd.AddCommand(launchCmd)

}

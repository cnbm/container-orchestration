package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launches the CNBM container orchestration benchmark",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Executing the CNBM container orchestration benchmark")
	},
}

func init() {
	RootCmd.AddCommand(launchCmd)

}

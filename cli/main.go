package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/cnbm/container-orchestration/cli/cmd"
)

func main() {
	if envd := os.Getenv("DEBUG"); envd != "" {
		log.SetLevel(log.DebugLevel)
	}
	cmd.Execute()
}

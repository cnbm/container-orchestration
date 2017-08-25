package dcos

import (
	"crypto/tls"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/cnbm/container-orchestration/pkg/generic"
	marathon "github.com/gambol99/go-marathon"
)

// Scalebench represents the DC/OS specific benchmark run for the scaling benchmark
type Scalebench struct {
	Config map[string]string
}

// Setup prepares and inits the DC/OS environment for the scaling benchmark
func (bench Scalebench) Setup() error {
	log.Info("Setting up DC/OS scaling benchmark")
	return nil
}

// Execute executes the scaling benchmark against a DC/OS cluster
func (bench Scalebench) Execute() (generic.Result, error) {
	log.Info("Executing DC/OS scaling benchmark")
	r := generic.Result{}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ingore unsigned cert
	}
	config := marathon.NewDefaultConfig()
	config.URL = bench.Config["dcosurl"]
	config.DCOSToken = bench.Config["dcosacstoken"]
	config.HTTPClient = &http.Client{Transport: tr}
	client, err := marathon.NewClient(config)
	if err != nil {
		return r, fmt.Errorf("Failed to create a client for marathon, error: %s", err)
	}
	// create
	// TODO prefetch
	log.Info("Deploying a new application")
	application := marathon.NewDockerApplication().
		Name("bench").
		CPU(0.1).
		Memory(64).
		Storage(0.0).
		Count(10).
		AddArgs("sleep 100000")

	application.
		Container.Docker.Container("ubuntu:xenial").
		Bridged().
		Expose(80).
		Expose(443)

	_, err = client.CreateApplication(application)
	if err != nil {
		return r, fmt.Errorf("Failed to create application: %s, error: %s", application, err)
	}
	log.Infof("Created the application: %s", application)

	// list applications
	applicationRunning, err := client.Application("bench")
	if err != nil {
		return r, fmt.Errorf("Failed to list application: %s", err)
	}
	log.Infof("Found %d instances running", applicationRunning.TasksRunning)
	return generic.Result{}, nil
}

// Teardown tears down and cleans up the DC/OS environment after the scaling benchmark has executed
func (bench Scalebench) Teardown() error {
	log.Info("Tearing down DC/OS scaling benchmark")
	return nil
}

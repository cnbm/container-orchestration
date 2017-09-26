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
type ApiCall struct {
	Config map[string]string
}

// Setup prepares and inits the DC/OS environment for the scaling benchmark
func (bench ApiCall) Setup() error {
	log.Info("Setting up DC/OS ApiCall benchmark")

	
	
	return nil
}

// Execute executes the distrubution benchmark against a DC/OS cluster
func (bench ApiCall) Execute() (generic.BenchmarkResult, error) {
	log.Info("Executing DC/OS ApiCall benchmark")
	r := generic.BenchmarkResult{}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ingore unsigned cert
	}
	config := marathon.NewDefaultConfig()
	config.URL = bench.Config["dcosurl"]
	config.DCOSToken = bench.Config["dcosacstoken"]
	config.HTTPClient = &http.Client{Transport: tr}
	client, err := marathon.NewClient(config)
	if err != nil {
		return r, fmt.Errorf("Failed to create a client for Marathon: %s", err)
	}
	// create
	// TODO prefetch
	log.Info("Deploying a new application")
	application := marathon.NewDockerApplication().
		Name("bench1").
		CPU(0.1).
		Memory(64).
		Storage(0.0).
		Count(1).
		AddArgs("/bin/sleep", "100000")

	application.
		Container.Docker.Container("busybox:1").
		Host().
		Expose(80)

	app, err := client.CreateApplication(application)
	if err != nil {
		return r, fmt.Errorf("Failed to create application %s: %s", application, err)
	}
	log.Infof("Creating the application: %s", application)

	// Wait for deployment with no timeout 
	err = client.WaitOnDeployment(app.Deployments[0]["id"], 0)
	if err != nil {
		return r, fmt.Errorf("Failed to list application: %s", err)
	}
	r.Output = fmt.Sprintf("Deployment complete.")
	return r, nil
}

// Teardown tears down and cleans up the DC/OS environment after the distribution benchmark has executed
func (bench ApiCall) Teardown() error {
	log.Info("Tearing down DC/OS ApiCall benchmark")
	tr := &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ingore unsigned cert
	}
	config := marathon.NewDefaultConfig()
	config.URL = bench.Config["dcosurl"]
	config.DCOSToken = bench.Config["dcosacstoken"]
	config.HTTPClient = &http.Client{Transport: tr}
	client, err := marathon.NewClient(config)
	if err != nil {
		return fmt.Errorf("Failed to create a client for Marathon: %s", err)
	}
	force := true
	_ , err = client.DeleteApplication("bench1", force)
	if err != nil {
		return fmt.Errorf("Failed to delte application: %s", err)
	}
	return nil
}

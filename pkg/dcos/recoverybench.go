package dcos

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os/exec"

	log "github.com/Sirupsen/logrus"
	"github.com/cnbm/container-orchestration/pkg/generic"
	marathon "github.com/gambol99/go-marathon"
)

// Scalebench represents the DC/OS specific benchmark run for the scaling benchmark
type Recovery struct {
	Config map[string]string
}

// Setup prepares and inits the DC/OS environment for the scaling benchmark
func (bench Recovery) Setup() error {
	// Start container on specific host.

	log.Info("Setting up DC/OS recovery benchmark")
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
	// create
	// TODO prefetch
	log.Info("Deploying a new application")
	application := new(marathon.Application).
		Name("bench1").
		CPU(0.5).
		Memory(64).
		Storage(0.0).
		Count(1).
		AddArgs("/bin/sleep", "100000")

	app, err := client.CreateApplication(application)
	if err != nil {
		return fmt.Errorf("Failed to create application %s: %s", application, err)
	}
	log.Infof("Creating the application: %s", application)

	// Wait for deployment with no timeout 
	err = client.WaitOnDeployment(app.Deployments[0]["id"], 0)
	if err != nil {
		return fmt.Errorf("Failed to list application: %s", err)
	}
	return nil
}

// Execute executes the recovery benchmark against a DC/OS cluster
func (bench Recovery) Execute() (generic.BenchmarkResult, error) {
	r := generic.BenchmarkResult{}

	// dcos task exec sleep  sh -c "ps axf | grep sleep | grep -v grep | awk '{print \"sudo kill -9 \" $1}' | sh"	
	cmd := exec.Command("dcos", "task","exec", "sleep", "sh", "-c", `"`, "ps axf | grep sleep | grep -v grep | awk '{print \"sudo kill -9 \" $1}' | sh", `"` )
	   // //""ps", "axf", "|", "grep", "sleep", "|", "grep",  "-v", "grep" , "|",  "awk",  "'{print",  "\\"sudo", "kill",  "-9", "\\"", "$1}'", "|", "sh\\"")
  err := cmd.Start()
  if err != nil {
  	log.Fatal(err)
  }
  log.Printf("Waiting for Recovery to finish...")
  err = cmd.Wait()

  //Check for deployment to finsih
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

	// TODO turn into while loop


  deployments, err := client.Deployments()
  if len(deployments) == 1 {
  	deployment := deployments[0]

		// Wait for deployment with no timeout 
		err = client.WaitOnDeployment(deployment.ID, 0)
		if err != nil {
			return r, fmt.Errorf("Failed to list application: %s", err)
		}
	}
	
	r.Output = fmt.Sprintf("Recovery complete.")
	return r, nil
}

// Teardown tears down and cleans up the DC/OS environment after the distribution benchmark has executed
func (bench Recovery) Teardown() error {
	log.Info("Tearing down DC/OS distribution benchmark")
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
		return fmt.Errorf("Failed to delete application: %s", err)
	}
	return nil
}

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
  "crypto/tls"
  "fmt"
  "log"
  "net/http"
  

  marathon "github.com/gambol99/go-marathon"
)

type scalebench struct {
  dcosUrl, dcosACSToken string 
}

func (bench scalebench) setup() error {
    fmt.Println("Setup")

    return nil
}

func (bench scalebench) execute() (result, error) {
    fmt.Println("Execute")

    // ingore unsigned cert
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    config := marathon.NewDefaultConfig()
    config.URL = bench.dcosUrl
    config.DCOSToken = bench.dcosACSToken
    config.HTTPClient = &http.Client{Transport: tr}
    client, err := marathon.NewClient(config)
    if err != nil {
      log.Fatalf("Failed to create a client for marathon, error: %s", err)
    }

    // create
    // TODO prefetch
    log.Printf("Deploying a new application")
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

    if _, err := client.CreateApplication(application); err != nil {
        log.Fatalf("Failed to create application: %s, error: %s", application, err)
    } else {
        log.Printf("Created the application: %s", application)
    }



    // list applications
    applicationRunning, err := client.Application("bench")

    fmt.Printf("Found %d instances running", applicationRunning.TasksRunning)


    return result{},nil
}

func (bench scalebench) teardown() error {
    fmt.Println("Teardown")
    return nil
}

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
  "fmt"
  "io/ioutil"
  "net/http"
  "crypto/tls"

  "github.com/tidwall/gjson"
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
    client := &http.Client{Transport: tr}

    header := make(http.Header)
    header.Set("Authorization", "token="+bench.dcosACSToken)
  
    req, _ := http.NewRequest("GET", bench.dcosUrl + "/marathon/v2/apps", nil)
    req.Header.Add("Authorization", "token="+bench.dcosACSToken)

    res, err := client.Do(req)

    if err != nil {
        panic(err)  
    }
    defer res.Body.Close()

    bodyBytes, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }

    appsJson := string(bodyBytes)
    fmt.Println(appsJson)

    //appName := "/display"
    instances := gjson.Get(appsJson, `apps.#[id="/display"].instances` ) //|  select(.id == "/display")| .instances
    
    fmt.Println("Instances running: ", instances.Int())

    return result{},nil
}

func (bench scalebench) teardown() error {
    fmt.Println("Teardown")
    return nil
}

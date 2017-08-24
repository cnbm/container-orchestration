// Copyright © 2017 The CNBM contributors
//
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

func main() {
	s := scalebench{
		dcosURL:      "https://joerg-22r-elasticl-mv9wyg0lclf4-218935880.us-west-2.elb.amazonaws.com",
		dcosACSToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1aWQiOiJib290c3RyYXB1c2VyIiwiZXhwIjoxNTAyMDg5NjAyfQ.gMxn7mw7om5tdsBJS6qtDzD6_mUR1ySzNndB0JS2ZiIUGOE6lYDFB2K22uWFi_hqKRg3RPHUdJI47esvY0DlWH20veLSDJEA9vRzg9qPLcKXzrCy_zwF_q1fw_uwkEIdVrmvttHmNEWiW4V1bbDajx9lWDFiKhz7d7p5BHaYvP1ycDhVTjDTBLyBAIC4CAdgFoh1MFocXfNk-SC4yXp68H4v13bTL7jhwjHpgeRWK_c2NH8J53vJUJdOuXqnTqNMKdbZ0D03kx5AlaNdMWpTiAMteschn9ZsdlaeihKLoqvPGPQR-emuOsua0h0njWobqAUhMVcsoere0eJF1_l5dg",
	}

	run(s)
}

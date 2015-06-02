// Copyright 2015 Bart van der Schans. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pingdomchecker

import (
	"log"
	"encoding/json"
	"os"
)

type Config struct {
	Username string
	Password string
	APIKey   string
}

// ClientConfig() returns a new client for accessing pingdom.
func ClientConfig(filename string) (Config) {

	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	config := Config{}

	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal("error:", err)
	}
	if  config.Username == "" {
		log.Fatal("Username is missing in config.")
	}
	if  config.Password == "" {
		log.Fatal("Password is missing in config.")
	}
	if  config.APIKey == "" {
		log.Fatal("APIKey is missing in config.")
	}

	return config
}

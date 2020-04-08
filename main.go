/*
Copyright © 2020 golark golark@outlook.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"github.com/golark/utask/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// init reads in viper config, sets up logrus
func init() {

	// step 1 - logrus config
	log.SetOutput(os.Stdout)     // use stdout rather than default stderr
	log.SetLevel(log.TraceLevel) // verbosity - warn and above

	// step 2 - viper config
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.SetConfigName("utaskcfg.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Can't read config file - falling back to default options")
	}

}

func main() {
	cmd.Execute()
}

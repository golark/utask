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

// Package cmd
// Cobra based CLI interface for user input
// Gets user request and passes the requests to cmdhandler
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "utask",
	Short: "utask is a methodology for tackling complex problems with focused tasks",
	Long: `utask cli tool is a companion for micro task methodology. 
This tool aims to help keep a high level of concentration while working on the task and to be mindful of taking breaks between workloads

usage:  $ utask start        # starts 25 minute utask but don't log to database
        $ utask start -t=30  # starts 30 minute utask but don't log to database'

        # start utask, log project, task name and notes
        $ utask start -p=<project name> -t=<task name> -n=<notes>  

Method: 
* Dissect complex projects into self contained bite-sized chunks of work called micro tasks
* Micro Tasks are focused bursts of incremental work to eventually lead to project completion
* Each Micro task can be registered with this tool to keep a log of all the micro tasks that contributed to the project `,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".utask" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".utask")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

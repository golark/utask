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

package cmd

import (
	"github.com/golark/utask/cmdhandler"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starts utask timer ",
	Long: `A single shot utask timer For example
     utask start            # starts default 25 minute utask timer
     utask start --timer=30 # starts 30 minute utask timer:
     utask start -d=30 -p=<project name> -t=<task name> -n=<notes> # starts 30 minute utask and logs the utask to database

     Minimum time interval is 1 minutes`,
	Run: func(cmd *cobra.Command, args []string) {

		// step 1 - get timer from flags - if no option is provided, default will be used
		sTimer, err := cmd.Flags().GetString("timer")
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("can not get timer flag - exiting")
			return
		}
		// convert time to integer
		iTimer, err := strconv.Atoi(sTimer)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("can not convert timer to integer - exiting")
			return
		}

		// step 2 - get project name from flags - if no option is provided, default will be used
		projName, err := cmd.Flags().GetString("projectname")
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("can not get a valid projectname - continue without")
		}

		// step 2 - get task name from flags - if no option is provided, default will be used
		taskName, err := cmd.Flags().GetString("taskname")
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("can not get a valid taskname- continue without")
		}

		// step 3 - get task note from flags - if no option is provided, default will be used
		taskNote, err := cmd.Flags().GetString("tasknote")
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("can not get a valid tasknote- continue without")
		}

		// trace the request
		log.WithFields(log.Fields{"time": iTimer, "name": taskName, "note": taskNote}).Trace("new utask")

		// step 4 - start single shot utask
		err = cmdhandler.Start(iTimer, projName, taskName, taskNote)
		if err != nil {
			log.Error("couldn't start the utask, please check input parameters and make sure daemon is running")
		}

	},
}

func init() {

	rootCmd.AddCommand(startCmd)

	// add flags
	startCmd.Flags().StringP("duration", "d", cmdhandler.DefaultTimeMins, "(optional) utask duration in minutes")
	startCmd.Flags().StringP("projectname", "p", "", "(optional) project name for utask")
	startCmd.Flags().StringP("taskname", "t", "utask", "(optional) task name for utask")
	startCmd.Flags().StringP("tasknote", "n", "keepitup", "(optional) task notes for utask")

}

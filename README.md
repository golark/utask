#![utasklogo](media/utasklogo_withprint.png)
 
 [![Go Report Card](https://goreportcard.com/badge/github.com/golark/utask)](https://goreportcard.com/report/github.com/golark/utask)

 
 CLI for the utask productivity tool that aims to boost focus by dissecting complex projects into small chunks of increments called utasks ( micro tasks ).
 
 ## About utask
 Feeling like doing work, all the while, staring at a screen with disarranged thoughs and random behaviour is a ubiquitous time waster.
 This seems to stem from four reasons, distraction, not having clear goals, getting stuck in complexity or diving deep into unnecessary details.
 The large amount of idle time can be optimised by focused chunks of systematic effort that dissolve complex projects into simple incremental tasks, that is what this project is about.
 
 - first step is dissecting projects into utaks ( microtasks ), which are self contained increments of work
 - utasks are tackled one at a time with highly focused 20-30 minute sessions in manner similar to the pomodoro technique
 - each utask is registered on a database for a post mortem after project completion
 
 Breaking down the project to simple self contained units of workload helps with timeline estimation to completion.
 Moreover, it provides a method to measure personal productivity levels. For example, this tool consisted of 23 utasks which were completed on a course of
 6 days, averaging 4 utasks per day. For future projects, effort estimations can be done in units of utasks, providing an accurate evaluation of the timeline. 
 
 ## utask cli

 This tool is the cli that sends users requests to the utask deamon. 
 It sends the requests over a Unix socket in form of Rest APIs, employing a similar architecture to Docker Cli <-> Docket daemon interface.
 Some of the cutting-edge libraries included in this package are Cobra, Viper, Logrus, Gorilla Mux.
 
 ## Installation
 
 Only manual installation is supported currently. Package manager installation will be available in the near future.
 
 ## Usage
 
 ` $ utask start -t=30 -n=<taskname> -m=<tasknotes> # starts 30 minute utask and registers the task name/notes`
 
 ` $ utask start                                    # starts default 25 minute utask timer`

 ` $ utask help                                     # further info about usage and commands`
 
 ## Contribution
 
 1. please check out tagged issues, and feature improvements
 2. as usual create a feature branch  `$ git checkout https://github.com/golark/utaskdaemon -b <featurebranch>`
 3. test and submit pull request ( feature -> develop )
 
 Following Gitflow guidelines. Code walkthrough can be arranged for the keen contributor. Any questions, please feel free to drop me a line.
 
 #### Architecture

 utask CLI handles the user interface for starting new utask entries and focus periods. Once a request is received, it parses
 the request parameters and communicates with the daemon. After the request is sent to the daemon, the cli immediately returns. 
 Daemon is then responsible to handle the user request. 
 
 utask CLI employs Cobra library as the front end interface to define potential commands, flags and arguments. The backend is handled
 by handlers that interface between the cli front end and daemon. The requests to the daemon are done via Rest API on a Unix Socket.

 ##### Directory structure and files
 
     .
     ├── cmd                         # cobra interface for cli commands
     |   ├── ping.go                 # ping daemong - mainly used for sanity checking
     |   ├── root.go                 # main root level command
     |   ├── start.go                # start utask timer command ( single shot )
     ├── cmdhandler                  # receives the command details from cobra and sends requests to daemon
     |   ├── ping.go                 # ping the daemon at the unix socket 
     |   ├── start.go                # daemon request to start single shot utask
     ├── media                       # images for project logo and installation demonstration
     |   ├── utasklogo.png           #
     |   ├── utasklogo_withprint.png #
     ├── config.yaml                 # configuration for the project  
     ├── go.mod                      # required modules
     ├── install.sh                  # for manual installations
     ├── LICENSE                     # 
     ├── main.go                     # entry point, configure cobra, logrus, viper, read config, handle signals
     ├── README.md                   #
     └ 
 ### Copyright and license
 
 Code and documentation released under the [Apache 2.0 License](LICENSE)
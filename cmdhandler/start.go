package cmdhandler

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"strconv"
)

const (
	MinTimeMins     int    = 1    // minimum time for a single shot utask
	DefaultTimeMins string = "25" // default time if no duration is specified during CLI
)

// Start a single shot utask timer by calling daemon Rest Api over unix socket
// pass timer duration, task name and note to the daemon
// return error if unsuccessful
func Start(iTimer int, taskName string, taskNote string) error {

	// step 1 - check that the time is not less than the minimum
	if iTimer < MinTimeMins {
		log.WithFields(log.Fields{"timer": iTimer, "minimum time": MinTimeMins}).Error("Time can not be less than minimum time")
		return errors.New("time can not be less than the minimum")
	}

	// step 2 - ping socket as a verification step
	err := Ping()
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("can not ping the socket")
		return err
	}

	// step 3 - connect to  the daemon over unixsocket
	httpc := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/tmp/utaskdaemon.sckt")
			},
		},
	}

	// step 4 - request daemon to star the timer
	_, err = httpc.Get("http://localhost/start/" + strconv.Itoa(iTimer) + "/" + taskName + "/" + taskNote)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("request to daemon returned err")
		return errors.New("can't connect to daemon")
	}

	return nil
}

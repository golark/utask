package cmdhandler

import (
	"context"
	log "github.com/sirupsen/logrus"
	errors "golang.org/x/xerrors"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

// Ping tries to ping the utask daemon over unix socket
// returns error if not successful
func Ping() error {

	// step 1 - create http client at the utask daemon socket
	httpc := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/tmp/utaskdaemon.sckt")
			},
		},
	}

	// step 2 - verify the existence of the socket
	fInfo, err := os.Stat("/tmp/utaskdaemon.sckt")
	if os.IsNotExist(err) {
		log.WithFields(log.Fields{"err": err}).Error("socket does not exist")
		return err
	}
	if fInfo.IsDir() { // a socket cannot be a directory, a dog is not a cat, a cat is not a dog, basics...
		log.WithFields(log.Fields{"err": err}).Error("socket seems to be a directory, aborting")
		return errors.New("socket can not be a directory, check daemon installation")
	}

	// step 3 - try to ping the unix socket and check for errors
	resp, err := httpc.Get("http://localhost/ping")
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error during converting time received over rest - using to default time")
		return err
	}
	defer resp.Body.Close()

	// step 4 - check status
	if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{"err": err}).Error("Ping didn't return StatusOK")
		return errors.New("Ping return code is not StatusOK")
	}

	// step 5 - log response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error when reading ping response body")
	} else {
		log.WithFields(log.Fields{"response body": string(body)}).Info("Ping Command returned")
	}

	return nil
}

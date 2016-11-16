package tekleader

import (
	"net/http"
	"time"
)

var Timeout int

func extendedTimeoutRequest() error {
	timeout := time.Duration(time.Duration(Timeout) * time.Second)
	netClient := &http.Client{
		Timeout: timeout,
	}
	_, err := netClient.Get(intraURL)
	if err != nil {
		return err
	}

	return nil
}

func getResRequest(req *http.Request) (*http.Response, error) {
	netClient := http.Client{}

	return netClient.Do(req)
}

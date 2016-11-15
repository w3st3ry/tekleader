package tekleader

import (
	"net/http"
	"time"

	"github.com/fatih/color"
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

func IntraIsAlive() {
	for {
		err := extendedTimeoutRequest()
		t := time.Now()
		if err != nil {
			color.Red("[%s] - Intranet is down... ./fixBocal.exe :noel:\n", t.Format(time.Stamp))
		} else {
			color.Green("[%s] - Intranet is alive :hap:\n\n", t.Format(time.Stamp))
			break
		}
	}
}

func getResRequest(req *http.Request) (*http.Response, error) {
	netClient := http.Client{}

	return netClient.Do(req)
}

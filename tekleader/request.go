package tekleader

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func disableTLS() *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

func extendedTimeoutRequest() error {
	timeout := time.Duration(time.Duration(Timeout) * time.Second)

	// Disable check cert of intra bc bocal suck

	netClient := &http.Client{
		Timeout:   timeout,
		Transport: disableTLS(),
	}

	_, err := netClient.Get(intraURL)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func getResRequest(req *http.Request) (*http.Response, error) {
	netClient := &http.Client{
		Transport: disableTLS(),
	}

	return netClient.Do(req)
}

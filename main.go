package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/w3st3ry/tekleader/cmd"
	"github.com/w3st3ry/tekleader/config"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Errorf("Could not load cfg file: %s", err)
	}
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

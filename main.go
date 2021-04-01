package main

import (
	"github.com/kostya2011/prometheus-ecs-sd/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

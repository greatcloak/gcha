package main

import (
	"github.com/greatcloak/gcha/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Great Cloak Hosted Apps CLI")

	cmd.Execute()
}

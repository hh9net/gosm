package main

import (
	"os"
	"os/signal"
	"syscall"

	"./checker"
	"./models"
	"./web"
)

var (
	checkChannel      = make(chan *models.Service)
	checkCountChannel chan (bool)
)

func main() {
	fixSIGTERM()
	models.CurrentConfig = models.ParseConfigFile()
	go web.Start()
	checker.Start()
}

func fixSIGTERM() {
	// Workaround for SIGTERM not working when pinging
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()
}

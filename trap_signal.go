package main

import (
	"os"
	"os/signal"
	"syscall"
)

func trapSignal() (os.Signal, error) {
	// Wait here until CTRL-C or other termnal signal is received.
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)
	return <-stopSignal, nil
}

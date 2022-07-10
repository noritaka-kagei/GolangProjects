package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	pid := os.Getpid()
	fmt.Printf("This program PID: %d\n", pid)

	// Wait here until CTRL-C or other termnal signal is received.
	stopSignal, err := trapSignal()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Received signal: %s\n", stopSignal)
}

func trapSignal() (os.Signal, error) {
	// Wait here until CTRL-C or other termnal signal is received.
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)
	return <-stopSignal, nil
}

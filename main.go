package main

import (
	"fmt"
	"os"
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

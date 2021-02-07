package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Akshit8/go-reminders/client"
)

var (
	backendURIFlag = flag.String("backend", "http://localhost:8000", "backend API uri")
	helpFlag       = flag.Bool("help", false, "display help manual for the CLI")
)

func main() {
	flag.Parse()

	s := client.NewSwitch(*backendURIFlag)

	if *helpFlag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("cmd switch error: %v\n", err)
		os.Exit(2)
	}
}

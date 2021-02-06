package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no command provided")
		os.Exit(3)
	}

	cmd := os.Args[1]
	switch cmd {
	case "greet":
		msg := "CLI BASICS"
		if len(os.Args) > 2 {
			f := strings.Split(os.Args[2], "=")
			if len(f) == 2 && f[0] == "--msg" {
				msg = f[1]
			}
		}
		fmt.Printf("greeting from cli: %s\n", msg)
	case "help":
		fmt.Printf("some help message\n")
	default:
		fmt.Printf("unknown command: %s\n", cmd)
	}
}

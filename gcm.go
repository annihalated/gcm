package main

import (
	"fmt"
	"gcm/cli"
	"os"
)

func main() {
	switch os.Args[1] {
	case "init":
		cli.Init()
	case "log":
		cli.Log()
	default:
		fmt.Println("Command not recognised")
	}
}

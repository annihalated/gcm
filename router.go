package main

import (
	"fmt"
	"gcm/cli"
	"os"
)

func Route() {
	if len(os.Args) == 1 {
		fmt.Println(welcome_message)
		os.Exit(3)
	}
	switch os.Args[1] {
	case "init":
		cli.Init()
	case "log":
		if val, _ := cli.Exists(".gcm/"); val == false {
			fmt.Println("Please initialise a gcm repository first")
			os.Exit(3)
		}

		if len(os.Args) <= 2 {
			fmt.Println("Please name your logged copy of the repository as an argument")
			fmt.Println("For example: ./gcm log v1")
			os.Exit(3)
		}

		cli.Log(os.Args[2])
	default:
		fmt.Println(command_not_valid_message)
	}
}

var welcome_message string = "Welcome to GCM! Please enter a command, like ./gcm init"
var command_not_valid_message string = "That wasn't a valid command"

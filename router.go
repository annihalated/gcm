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
	case "snap":
		if val, _ := cli.PathExists(".gcm/"); val == false {
			fmt.Println("Please initialise a gcm repository first")
			os.Exit(3)
		}

		if len(os.Args) <= 2 {
			fmt.Println("Please name your snapshot of the repository as an argument")
			fmt.Println("For example: ./gcm snap v1")
			os.Exit(3)
		}

		cli.MakeSnapshot(os.Args[2])

	case "log":
		if val, _ := cli.PathExists(".gcm/"); val == false {
			fmt.Println("Please initialise a gcm repository first")
			os.Exit(3)
		}

		cli.DisplayLog()
	default:
		fmt.Println(command_not_valid_message)
	}
}

var welcome_message string = "Welcome to GCM! Please enter a command, like ./gcm init"
var command_not_valid_message string = "That wasn't a valid command"

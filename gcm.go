package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	Route()
}

func Route() {
	if len(os.Args) == 1 {
		fmt.Println(welcome_message)
		os.Exit(3)
	}
	switch os.Args[1] {
	case "init":
		os.Mkdir(".gcm", os.ModePerm)
		os.Mkdir(".gcm/snapshots/", os.ModePerm)
		os.Create(INDEX_PATH)
		os.Create(".gcm/HEAD")
		fmt.Println("Initialized gcm repository in this directory")

	case "snap":
		checkForInit()
		if len(os.Args) <= 1 {
			fmt.Println("Snapshot names are autogenerated")
			os.Exit(3)
		}
		MakeSnapshot()

	case "log":
		checkForInit()
		DisplayLog()

	case "switch":
		checkForInit()
		if len(os.Args) != 3 {
			fmt.Printf("You have the wrong number of arguments.")
			os.Exit(3)
		}
		SwitchHEAD(os.Args[2])

	case "diff":
		checkForInit()
		Diff(os.Args[2], os.Args[3])
	case "head":
		checkForInit()
		fmt.Printf("%s", HEAD())

	default:
		fmt.Println(command_not_valid_message)
	}
}

var welcome_message string = "Welcome to GCM! Please enter a command, like ./gcm init"
var command_not_valid_message string = "That wasn't a valid command"

func checkForInit() (bool, error) {
	_, err := os.Stat(INDEX_PATH)
	if err != nil {
		log.Fatal("Please run the init command first.")
		os.Exit(3)
	}
	return true, nil
}

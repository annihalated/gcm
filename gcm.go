package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "init":
		fmt.Println("Initialized gcm repository in this directory")
		os.Mkdir(".gcm", os.ModePerm)
		os.Create(".gcm/log")
	case "log":
		if val, _ := exists(".gcm/"); val == true {
			if len(os.Args) > 2 {
				var logname string = ".gcm/" + os.Args[2]
				os.Mkdir(logname, os.ModePerm)
			} else {
				fmt.Println("Please name your logged copy of the repository as an argument")
				fmt.Println("For example: ./gcm log v1")
			}
		} else {
			fmt.Println("Please initialise a gcm repository first")
		}
	default:
		fmt.Println("Command not recognised")
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

package cli

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Log() {
	if val, _ := exists(".gcm/"); val == true {
		if len(os.Args) > 2 {

		} else {
			fmt.Println("Please name your logged copy of the repository as an argument")
			fmt.Println("For example: ./gcm log v1")
		}
	} else {
		fmt.Println("Please initialise a gcm repository first")
	}
}

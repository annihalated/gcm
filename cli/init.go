package cli

import (
	"fmt"
	"os"
)

func Init() {
	os.Mkdir(".gcm", os.ModePerm)
	os.Create(".gcm/log")

	fmt.Println("Initialized gcm repository in this directory")
}

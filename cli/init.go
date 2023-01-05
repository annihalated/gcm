package cli

import (
	"fmt"
	"os"
)

func Init() {
	os.Mkdir(".gcm", os.ModePerm)
	os.Create(".gcm/gcm.json")
	fmt.Println("Initialized gcm repository in this directory")
}

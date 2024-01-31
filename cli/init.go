package cli

import (
	"fmt"
	"os"
)

func Init() {
	os.Mkdir(".gcm", os.ModePerm)
	os.Mkdir(".gcm/snapshots/", os.ModePerm)
	os.Create(".gcm/gcm.json")
	os.Create(".gcm/HEAD")
	fmt.Println("Initialized gcm repository in this directory")
}

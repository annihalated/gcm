package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func HEAD() string {
	b, err := os.ReadFile(".gcm/HEAD")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	if str != "" {
		err = uuid.Validate(str)
		if err != nil {
			panic(err)
		}
		return str
	}

	return string("BASE")
}

func SwitchHEAD(direction string) bool {
	jsonFile, _ := os.ReadFile(".gcm/gcm.json")
	json.Unmarshal(jsonFile, &snapshots)

	if direction == HEAD() {
		fmt.Printf("HEAD is already at %s", direction)
		return false
	}
	hsc := HEAD()

	for _, snapshot := range snapshots {
		if snapshot.Name == direction {
			_ = os.WriteFile(".gcm/HEAD", []byte(snapshot.Name), 0644)
			fmt.Printf("HEAD changed: %s -> %s\n", hsc, direction)
			return true
		}
	}

	fmt.Printf("%s is not a valid commit", direction)
	return false
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func SwitchHEAD(direction string) bool {
	jsonFile, _ := os.ReadFile(".gcm/gcm.json")
	json.Unmarshal(jsonFile, &snapshots)

	if direction == GetHEADString() {
		fmt.Printf("HEAD is already at %s", direction)
		return false
	}
	hsc := GetHEADString()

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

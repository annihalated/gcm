package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func DisplayLog() bool {
	jsonFile, _ := os.ReadFile(".gcm/gcm.json")
	json.Unmarshal(jsonFile, &snapshots)

	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	for _, snapshot := range snapshots {
		fmt.Printf("%s - %s - from %s\n", red(snapshot.Time), green(snapshot.Name), blue(snapshot.Parent))
	}

	fmt.Printf("\n")

	return true
}

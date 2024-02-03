package main

import (
	"fmt"

	"github.com/fatih/color"
)

func DisplayLog() bool {
	snapshots := getSnapshots()

	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	for _, snapshot := range snapshots {
		fmt.Printf("%s - %s - from %s\n", red(snapshot.Time), green(snapshot.Name), blue(snapshot.Parent))
	}

	fmt.Printf("\n")

	return true
}

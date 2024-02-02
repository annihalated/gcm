package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func PrettySnapshot(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func DisplayLog() bool {
	jsonFile, _ := os.ReadFile(".gcm/gcm.json")
	json.Unmarshal(jsonFile, &snapshots)

	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	for _, snapshot := range snapshots {
		fmt.Printf("%s - %s - from %s\n", red(snapshot.Time), green(snapshot.Name), blue(snapshot.Parent))
	}

	return true
}

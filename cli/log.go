package cli

import (
	"encoding/json"
	"fmt"
	"os"
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

	for _, snapshot := range snapshots {
		fmt.Println(snapshot.Name, " - ", snapshot.Time, " - ", snapshot.Paths)
	}

	return true
}

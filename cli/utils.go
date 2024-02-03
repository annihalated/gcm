package cli

import (
	"encoding/json"
	"os"
	"strings"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Remove(s []string, r string) []string {
	var newSlice []string
	for _, v := range s {
		if !strings.Contains(v, r) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func PrettySnapshot(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

package main

import (
	"strings"
)

func Remove(s []string, r string) []string {
	var newSlice []string
	for _, v := range s {
		if !strings.Contains(v, r) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// Diff uses sergi/go-diff to produce a diff between files that are present in two snapshots.
func Diff(from string, to string) bool {
	var diffSnapshots []Snapshot
	snapshots := getSnapshots()
	for _, snapshot := range snapshots {
		if snapshot.Name == from {
			fmt.Println("Your FROM snapshot is " + snapshot.Name)
			diffSnapshots = append(diffSnapshots, snapshot)
		} else if snapshot.Name == to {
			fmt.Println("Your TO snapshot is " + snapshot.Name)
			diffSnapshots = append(diffSnapshots, snapshot)
		}
	}

	for _, snapshot := range diffSnapshots {
		fmt.Printf("%s: %s\n", snapshot.Name, snapshot.Paths)
	}

	paths1 := diffSnapshots[0].Paths
	paths2 := diffSnapshots[1].Paths

	for _, path1 := range paths1 {
		fmt.Println("We have the original file: " + path1)
		for _, path2 := range paths2 {
			if strings.EqualFold(path1, path2) == true {
				fmt.Println("We found the corresponding file: " + path2)
				fmt.Println("Here's a diff of both files: ")
				filepath1 := filepath.Join(".gcm", "snapshots", diffSnapshots[0].Name, path1)
				filepath2 := filepath.Join(".gcm", "snapshots", diffSnapshots[1].Name, path2)
				fmt.Println(filepath1)
				fmt.Println(filepath2)
				text1, err := os.ReadFile(filepath1)
				text2, err := os.ReadFile(filepath2)

				if err != nil {
					fmt.Println(err)
				}
				if strings.EqualFold(string(text1), string(text2)) != true {
					if utf8.ValidString(string(text1)) && utf8.ValidString(string(text2)) {
						dmp := diffmatchpatch.New()
						diffs := dmp.DiffMain(string(text1), string(text2), true)

						fmt.Println(dmp.DiffPrettyText(diffs))
					}
				}

			}

		}

	}

	return true

}

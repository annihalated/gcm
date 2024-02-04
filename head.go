package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

// HEAD() gets the current HEAD value from the file
// and checks that it is a valid UUID.
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
	snapshots := getSnapshots()
	if direction == HEAD() {
		fmt.Printf("HEAD is already at %s", direction)
		return false
	}
	hsc := HEAD()

	for _, snapshot := range snapshots {
		if snapshot.Name == direction {
			_ = os.WriteFile(".gcm/HEAD", []byte(snapshot.Name), 0644)
			fmt.Printf("HEAD changed: %s -> %s\n", hsc, direction)
			ReconstructDirTree(direction)
			return true
		}
	}

	fmt.Printf("%s is not a valid commit", direction)
	return false
}

func ReconstructDirTree(snapshotName string) (bool, error) {
	wd_index, err := CreateIndex(".")
	if err != nil {
		log.Fatal(err)
	}
	path_to_snapshot := ".gcm/snapshots/" + snapshotName + "/"
	snapshot_index, err := CreateIndex(path_to_snapshot)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", wd_index)
	fmt.Printf("%s", snapshot_index)

	for _, path := range wd_index {
		err = os.RemoveAll(path)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, path := range snapshot_index {
		fileInfo, err := os.Stat(path)
		path_without_prefix := strings.ReplaceAll(path, path_to_snapshot, "")
		if err != nil {
			log.Fatal(err)
		}
		if fileInfo.IsDir() {
			os.Mkdir(path, os.ModePerm)
		} else {
			copy(path, "./"+path_without_prefix)
		}
	}

	return true, nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func copy(src string, dst string) {
	// Read all content of src to data, may cause OOM for a large file.
	data, err := os.ReadFile(src)
	checkErr(err)
	// Write data to dst
	err = os.WriteFile(dst, data, 0644)
	checkErr(err)
}

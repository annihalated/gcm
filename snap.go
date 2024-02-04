package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Snapshot struct {
	Name   string
	Paths  []string
	Time   string
	Parent string
}

// CreateIndex returns a slice after recursively walking through
// the directory tree at and below a given path. It omits the ".gcm" and
// the ".git" paths.
func CreateIndex(path string) ([]string, error) {
	var paths []string
	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		paths = append(paths, path)
		return nil
	})
	paths = slices.DeleteFunc(paths, func(s string) bool {
		if strings.Contains(s, ".gcm") || strings.Contains(s, ".git") {
			return true
		}
		return false
	})
	paths = paths[1:]
	return paths, nil
}

func MakeSnapshot() bool {
	paths, err := CreateIndex(".")
	if err != nil {
		log.Fatal(err)
	}

	snapshotName := uuid.NewString()

	for _, path := range paths {
		stat, err := os.Stat(path)
		if err != nil {
			log.Fatal(err)
		}

		srcPath := path
		dstPath := filepath.Join(".gcm/snapshots/"+snapshotName+"/", srcPath)

		srcFile, err := os.Open(srcPath)
		if err != nil {
			log.Fatal(err)
		}

		defer srcFile.Close()

		if stat.IsDir() {
			err = os.Mkdir(dstPath, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err = os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}

			dstFile, err := os.Create(dstPath)
			if err != nil {
				log.Fatal(err)
			}

			_, err = io.Copy(dstFile, srcFile)
			if err != nil {
				log.Fatal(err)
			}

		}
	}

	t := time.Now().Format("Mon Jan 2 15:04:05 MST 2006")
	AppendSnapshotLog(snapshotName, paths, t, HEAD())
	fmt.Printf("Current version saved as %s", snapshotName)
	return true
}

func AppendSnapshotLog(snapshotName string, paths []string, t string, parent string) {
	snapshots := getSnapshots()
	snapshots = append(snapshots, Snapshot{
		Name:   snapshotName,
		Paths:  paths,
		Time:   t,
		Parent: HEAD(),
	})

	data, err := json.Marshal(snapshots)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(INDEX_PATH, data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(".gcm/HEAD", []byte(snapshotName), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getSnapshots() []Snapshot {
	var snapshots []Snapshot

	jsonFile, err := os.Open(INDEX_PATH)
	if err != nil {
		log.Fatal(err)
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	if string(byteValue) == "" {
		return []Snapshot{}
	}

	err = json.Unmarshal(byteValue, &snapshots)
	if err != nil {
		log.Fatal(err)
	}

	return snapshots
}

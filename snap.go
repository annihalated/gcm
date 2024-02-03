package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
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
func CreateIndex(path string) []string {
	var paths []string
	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		paths = append(paths, path)
		return nil
	})
	paths = Remove(paths, ".gcm")
	paths = Remove(paths, ".git")
	paths = paths[1:]
	return paths
}

func MakeSnapshot() bool {
	paths := CreateIndex(".")
	snapshotName := uuid.NewString()

	for _, path := range paths {
		stat, _ := os.Stat(path)
		srcPath := path
		dstPath := filepath.Join(".gcm/snapshots/"+snapshotName+"/", srcPath)

		srcFile, _ := os.Open(srcPath)

		defer srcFile.Close()

		if stat.IsDir() {
			os.Mkdir(dstPath, os.ModePerm)
		} else {
			os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
			dstFile, _ := os.Create(dstPath)
			io.Copy(dstFile, srcFile)
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
	data, _ := json.Marshal(snapshots)
	_ = os.WriteFile(".gcm/gcm.json", data, 0644)
	_ = os.WriteFile(".gcm/HEAD", []byte(snapshotName), 0644)
}

func PrettySnapshot(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func getSnapshots() []Snapshot {
	var snapshots []Snapshot
	jsonFile, _ := os.Open(".gcm/gcm.json")
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &snapshots)
	return snapshots
}

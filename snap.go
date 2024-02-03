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

var paths []string
var snapshots []Snapshot

func MakeSnapshot() bool {
	snapshotName := uuid.NewString()
	filepath.WalkDir(".", visit)
	paths = Remove(paths, ".gcm")
	paths = Remove(paths, ".git")
	paths = paths[1:]

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

	fmt.Printf("Current version saved as %s", snapshotName)

	t := time.Now()
	_, byteValue := ReadLog()
	json.Unmarshal(byteValue, &snapshots)
	AppendSnapshotLog(snapshotName, t)
	return true
}

func AppendSnapshotLog(snapshotName string, t time.Time) {
	formattedTime := t.Format("Mon Jan 2 15:04:05 MST 2006")
	snapshots = append(snapshots, Snapshot{
		Name:   snapshotName,
		Paths:  paths,
		Time:   formattedTime,
		Parent: HEAD(),
	})
	data, _ := json.Marshal(snapshots)
	_ = os.WriteFile(".gcm/gcm.json", data, 0644)
	_ = os.WriteFile(".gcm/HEAD", []byte(snapshotName), 0644)
}

func ReadLog() (error, []byte) {
	jsonFile, _ := os.Open(".gcm/gcm.json")
	byteValue, _ := io.ReadAll(jsonFile)
	return nil, byteValue
}

func visit(path string, di fs.DirEntry, err error) error {
	paths = append(paths, path)
	return nil
}

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var paths []string
var snapshots []Snapshot

func MakeSnapshot(snapshot_name string) bool {
	filepath.WalkDir(".", visit)
	paths = Remove(paths, ".gcm")
	paths = Remove(paths, ".git")
	paths = paths[1:]

	for _, path := range paths {
		stat, _ := os.Stat(path)
		srcPath := path
		dstPath := filepath.Join(".gcm/snapshots/"+snapshot_name+"/", srcPath)

		srcFile, _ := os.Open(srcPath)

		defer srcFile.Close()

		if stat.IsDir() {
			os.Mkdir(dstPath, os.ModePerm)
		} else {
			os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
			dstFile, _ := os.Create(dstPath)
			defer dstFile.Close()
			io.Copy(dstFile, srcFile)

		}
	}

	fmt.Printf("Current version saved as %s", snapshot_name)

	t := time.Now()
	jsonFile, _ := os.Open(".gcm/gcm.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &snapshots)
	snapshots = append(snapshots, Snapshot{
		Name:  snapshot_name,
		Paths: paths,
		Time:  t.String(),
	})
	data, _ := json.Marshal(snapshots)
	_ = ioutil.WriteFile(".gcm/gcm.json", data, 0644)
	return true
}

func visit(path string, di fs.DirEntry, err error) error {
	paths = append(paths, path)
	return nil
}

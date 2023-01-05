package cli

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var paths []string

func Log(log_name string) bool {
	paths := get_paths()
	destination := ".gcm/" + log_name + "/"
	duplicate(paths, destination)
	path_string := strings.Join(paths, ", ")
	t := time.Now()
	log_update := t.String() + " - " + log_name + ": " + path_string + "\n"
	UpdateLogFile(log_update)
	fmt.Printf("Current version saved as %s", log_update)
	return true
}

func get_paths() []string {
	filepath.WalkDir(".", visit)
	paths = Remove(paths, ".gcm")
	paths = Remove(paths, ".git")
	paths = paths[1:]
	return paths
}

func visit(path string, di fs.DirEntry, err error) error {
	paths = append(paths, path)
	return nil
}

func duplicate(paths []string, destination string) (bool, error) {
	for _, path := range paths {
		stat, _ := os.Stat(path)
		srcPath := path
		dstPath := filepath.Join(destination, srcPath)

		srcFile, _ := os.Open(srcPath)

		defer srcFile.Close()

		if stat.IsDir() == true {
			os.Mkdir(dstPath, os.ModePerm)
		} else {
			os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
			dstFile, _ := os.Create(dstPath)
			defer dstFile.Close()
			io.Copy(dstFile, srcFile)

		}
	}

	return true, nil
}

func UpdateLogFile(log string) error {
	f, err := os.OpenFile(".gcm/log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(log); err != nil {
		fmt.Println(err)
	}

	return nil
}

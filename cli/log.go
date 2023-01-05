package cli

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func Log(log_name string) bool {
	paths := get_paths()
	fmt.Printf("%v", paths)
	destination := ".gcm" + "/" + log_name + "/"
	duplicate(paths, destination)
	return true
}

var paths []string

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

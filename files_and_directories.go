package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func WalkFilesInDirectory(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func FilesInDirectory(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		if !file.IsDir() {
			file_path := fmt.Sprintf("%v/%v", root, file.Name())
			files = append(files, file_path)
		}
	}
	return files, nil
}

func FileExists(file_name string) bool {
	if _, err := os.Stat(file_name); err == nil {
		return true
	}
	return false
}

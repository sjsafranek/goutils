package utils

import (
    "os"
    "fmt"
    "path/filepath"
    "io/ioutil"
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

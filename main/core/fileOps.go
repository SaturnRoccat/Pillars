package core

import (
	"os"
	"path/filepath"
)

func GetFilesInDir(dir string) []string {
	// Get all files in the dir and return them in a slice
	var files []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	return files
}

func BulkFileMove(source, target string, filesToIgnore map[string]bool) error {
	// Recursively find all files in the source dir and check if they are in the filesToIgnore map
	// If they are not, move them to the target dir
	// If they are, do nothing
	// If there is an error, return false

	// Get all files in the source dir
	files := GetFilesInDir(source)

	// Loop through all files
	for _, file := range files {
		// Check if the file is in the filesToIgnore map
		if _, ok := filesToIgnore[file]; !ok {
			// Move the file to the target dir
			err := os.Rename(file, target)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

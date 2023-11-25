package PC

import (
	"io"
	"os"
	"path/filepath"
)

func GetFilesInDir(dir string) []string {
	// Get all files in the dir and return them in a slice
	var files []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files
}

func GetFilesInDirWithExt(dir string, ext string) []string {
	// Get all files in the dir and return them in a slice
	var files []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(path) != ext {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files
}

func GetFileInfoInDir(dir string) map[string]os.FileInfo {
	// Get all files in the dir and return them in a map
	var files = make(map[string]os.FileInfo)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files[path] = info
		return nil
	})
	return files
}

func CopyFile(source, destination string) error {
	// Open the source file
	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file
	var _, fileExt string = filepath.Split(source)

	destFile, err := os.Create(destination + "\\" + fileExt)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	// Flush any buffered data to ensure the file is written completely
	err = destFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func BulkFileCopy(source, target string, filesToIgnore map[string]bool) error {
	// Recursively find all files in the source dir and check if they are in the filesToIgnore map
	// If they are not, move them to the target dir
	// If they are, do nothing
	// If there is an error, return false

	// Get all files in the source dir
	files := GetFilesInDir(source)

	// Loop through all files
	for _, file := range files {
		var _, fileExt = filepath.Split(file)
		// Check if the file is in the filesToIgnore map
		if _, ok := filesToIgnore[fileExt]; !ok {
			// Move the file to the target dir
			err := CopyFile(file, target)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func DeleteContents(directory string) error {
	dirEntries, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, entry := range dirEntries {
		entryPath := filepath.Join(directory, entry.Name())

		if entry.IsDir() {
			err := os.RemoveAll(entryPath) // RemoveAll deletes files and subdirectories recursively
			if err != nil {
				return err
			}
		} else {
			err := os.Remove(entryPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

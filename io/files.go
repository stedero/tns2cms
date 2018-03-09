package io

import (
	"io/ioutil"
	"os"
)

// IsExistingDirectory determines whether a given directory
// is an existing directory.
func IsExistingDirectory(dir string) bool {
	file, err := os.Stat(dir)
	return err == nil && file.IsDir()
}

// CreateDirIfNotExist creates a directory if it does not
// exist yet. If an error occurs then the program terminates
// with a panic message.
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		checkError(os.MkdirAll(dir, 0755))
	}
}

// SelectFiles selects files from the given directory filtered by the
// given accept function. If an error occurs then the program
// terminates with a panic message.
func SelectFiles(dirname string, accept func(file os.FileInfo) bool) []os.FileInfo {
	var selected []os.FileInfo
	allFiles := readDir(dirname)
	for _, file := range allFiles {
		if accept(file) {
			selected = append(selected, file)
		}
	}
	return selected
}

// ReadFile reads an entire file into memory. If an error occurs then
// the program terminates with a panic message.
func ReadFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	checkError(err)
	return data
}

// WriteFile writes data to a file named by filename.
// If the file does not exist, WriteFile creates it with permissions perm;
// otherwise WriteFile truncates it before writing. If an error occurs then
// the program terminates with a panic message.
func WriteFile(filename string, data []byte) {
	checkError(ioutil.WriteFile(filename, data, 0644))
}

func readDir(dirname string) []os.FileInfo {
	files, err := ioutil.ReadDir(dirname)
	checkError(err)
	return files
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

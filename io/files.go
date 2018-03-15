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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

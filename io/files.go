package io

import (
	"io/ioutil"
	"log"
	"os"
)

// TnsXML info struct.
type TnsXML struct {
	FileName string
	Data     []byte
}

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
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatalf("fail to create directory %s: %v", dir, err)
		}
	}
}

// ReadFile reads an entire file into memory. If an error occurs then
// the program terminates with a panic message.
func ReadFile(filename string) *TnsXML {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("fail to read file %s: %v", filename, err)
	}
	return &TnsXML{filename, data}
}

// WriteFile writes data to a file named by filename.
// If the file does not exist, WriteFile creates it with permissions perm;
// otherwise WriteFile truncates it before writing. If an error occurs then
// the program terminates with a panic message.
func WriteFile(filename string, data []byte) {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatalf("fail to write file %s: %v", filename, err)
	}
}

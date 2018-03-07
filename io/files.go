package io

import (
	"io/ioutil"
	"os"
)

func IsExistingDirectory(dir string) bool {
	file, err := os.Stat(dir)
	return err == nil && file.IsDir()
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		checkError(os.MkdirAll(dir, 0755))
	}
}

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

func ReadFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	checkError(err)
	return data
}

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

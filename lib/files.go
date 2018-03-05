package lib

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
		err = os.MkdirAll(dir, 0755)
		checkError(err)
	}
}

func ReadDir(dirname string) (files []os.FileInfo) {
	files, err := ioutil.ReadDir(dirname)
	checkError(err)
	return files
}

func ReadFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	checkError(err)
	return data
}

func WriteFile(filename string, data []byte) {
	checkError(ioutil.WriteFile(filename, data, 0644))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

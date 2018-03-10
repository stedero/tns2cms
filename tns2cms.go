// tns2cms adds metadata to TNS articles for bulkload in the Alfresco CMS.
// Every TNS article is stored in a separate XML file.
// Metadata is created by extracting and converting data from a TNS article
// and store that in a separate XML file that complies with the JAVA properties
// file DTD (http://java.sun.com/dtd/properties.dtd).
package main

import (
	"os"
	"sync"
	"tns2cms/cmd"
	"tns2cms/io"
	"tns2cms/model"
	"tns2cms/naming"
	"tns2cms/stats"
)

// Process files in chunks of 100 to avoid error "too much open files"
const chunkSize = 100

func main() {
	var waitGroup sync.WaitGroup
	directoryNamer := naming.NewDirectoryNamer(cmd.ParseCommandLine())
	io.CreateDirIfNotExist(directoryNamer.OutDir())
	allFiles := io.SelectFiles(directoryNamer.InDir(), naming.Accept)
	defer stats.Reporter(len(allFiles))()
	nextFile := stats.ProgressIndicator(len(allFiles))
	for _, chunk := range chunks(chunkSize, allFiles) {
		waitGroup.Add(len(chunk))
		for _, file := range chunk {
			go func(fifo os.FileInfo) {
				processFile(directoryNamer.NewFilenamer(fifo))
				nextFile()
				waitGroup.Done()
			}(file)
		}
		waitGroup.Wait()
	}
}

func processFile(fileNamer *naming.Filenamer) {
	tnsXML := io.ReadFile(fileNamer.InputFilename())
	io.WriteFile(fileNamer.OutputFilename(), tnsXML)
	tnsArticle := model.NewTnsArticle(tnsXML)
	metaXML := model.NewMetaData(tnsArticle)
	io.WriteFile(fileNamer.MetaFilename(), metaXML)
}

func chunks(chunksize int, files []os.FileInfo) [][]os.FileInfo {
	rows := len(files)/chunksize + 1
	chunks := make([][]os.FileInfo, rows)
	for i := range chunks {
		start := i * chunksize
		end := start + min(chunksize, len(files)-start)
		chunks[i] = files[start:end]
	}
	return chunks
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

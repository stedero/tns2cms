// tns2cms adds metadata to TNS articles for bulkload in the Alfresco CMS.
// Every TNS article is stored in a separate XML file.
// Metadata is created by extracting and converting data from a TNS article
// and store that in a separate XML file that complies with the JAVA properties
// file DTD (http://java.sun.com/dtd/properties.dtd).
package main

import (
	"tns2cms/lib"
)

func main() {
	directoryNamer := lib.ParseCommandLine()
	lib.CreateDirIfNotExist(directoryNamer.OutDir())
	files := lib.ReadDir(directoryNamer.InDir())
	counter := lib.NewCounter(len(files))
	for _, file := range files {
		if lib.Accept(file) {
			processFile(directoryNamer.NewFilenamer(file))
		}
		counter.Next()
	}
}

func processFile(fileNamer *lib.Filenamer) {
	tnsXML := lib.ReadFile(fileNamer.InputFilename())
	lib.WriteFile(fileNamer.OutputFilename(), tnsXML)
	tnsArticle := lib.NewTnsArticle(tnsXML)
	metaXML := lib.NewMetaData(tnsArticle)
	lib.WriteFile(fileNamer.MetaFilename(), metaXML)
}

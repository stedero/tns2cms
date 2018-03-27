// tns2cms adds metadata to TNS articles for bulkload in the Alfresco CMS.
// Every TNS article is stored in a separate XML file.
// Metadata is created by extracting and converting data from a TNS article
// and store that in a separate XML file that complies with the JAVA properties
// file DTD (http://java.sun.com/dtd/properties.dtd).
package main

import (
	"ibfd.org/tns2cms/cmd"
	"ibfd.org/tns2cms/exec"
	"ibfd.org/tns2cms/io"
	"ibfd.org/tns2cms/model"
	"ibfd.org/tns2cms/paths"
	"ibfd.org/tns2cms/stats"
)

func main() {
	verbose, directoryNamer := cmd.ParseCommandLine()
	reporter := stats.NewReporter(verbose)
	visitor := exec.NewVisitor(directoryNamer, process, reporter)
	visitor.Walk()
	reporter.End()
}

func process(fileNamer *paths.Filenamer) {
	tnsXML := io.ReadFile(fileNamer.InputFilename())
	io.WriteFile(fileNamer.OutputFilename(), tnsXML)
	tnsArticle := model.NewTnsArticle(tnsXML)
	metaXML := model.NewMetaData(tnsArticle)
	io.WriteFile(fileNamer.MetaFilename(), metaXML)
}

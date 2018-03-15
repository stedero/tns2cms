package exec

import (
	"os"
	"path/filepath"
	"tns2cms/io"
	"tns2cms/naming"
	"tns2cms/stats"
)

// Visitor defines a visitor
type Visitor struct {
	rootDirNamer *naming.DirectoryNamer
	process      func(Filenamer *naming.Filenamer)
	reporter     *stats.Reporter
}

// NewVisitor creates a new visitor
func NewVisitor(rootDirNamer *naming.DirectoryNamer, processor func(*naming.Filenamer), reporter *stats.Reporter) *Visitor {
	return &Visitor{rootDirNamer, processor, reporter}
}

// Walk the directory tree and process each file
func (visitor *Visitor) Walk() error {
	return filepath.Walk(visitor.rootDirNamer.InDir(), visitor.walker())
}

// Walker returns a directory walker function.
func (visitor *Visitor) walker() func(string, os.FileInfo, error) error {
	rootDirNamer := visitor.rootDirNamer
	reporter := visitor.reporter
	return func(path string, fileInfo os.FileInfo, err error) error {
		validation := naming.Validate(fileInfo)
		defer reporter.Register(validation, path)
		switch validation {
		case naming.AcceptFile:
			fileNamer := naming.NewFilenamerFromRoot(rootDirNamer, path, fileInfo)
			visitor.process(fileNamer)
		case naming.AcceptDir:
			dest := rootDirNamer.NewOutdirName(path)
			io.CreateDirIfNotExist(dest)
		case naming.RejectDir:
			return filepath.SkipDir
		case naming.RejectFile:
		default:
		}
		return err
	}
}

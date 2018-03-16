package exec

import (
	"os"
	"path/filepath"
	"tns2cms/io"
	"tns2cms/paths"
	"tns2cms/stats"
)

// Visitor defines a visitor
type Visitor struct {
	rootDirNamer *paths.DirectoryNamer
	process      func(Filenamer *paths.Filenamer)
	reporter     *stats.Reporter
}

// NewVisitor creates a new visitor
func NewVisitor(rootDirNamer *paths.DirectoryNamer, processor func(*paths.Filenamer), reporter *stats.Reporter) *Visitor {
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
		validation := paths.Validate(fileInfo)
		defer reporter.Register(validation, path)
		switch validation {
		case paths.AcceptFile:
			fileNamer := paths.NewFilenamerFromRoot(rootDirNamer, path, fileInfo)
			visitor.process(fileNamer)
		case paths.AcceptDir:
			dest := rootDirNamer.NewOutdirName(path)
			io.CreateDirIfNotExist(dest)
		case paths.RejectDir:
			return filepath.SkipDir
		case paths.RejectFile:
		default:
		}
		return err
	}
}

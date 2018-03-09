package naming

import (
	"os"
	"path/filepath"
	"strings"
)

const extension = ".xml"
const metaFileSuffix = ".metadata.properties" + extension

// Filenamer holds all data needed to create input- and outputfilenames.
type Filenamer struct {
	indir  string
	outdir string
	file   os.FileInfo
}

// InputFilename returns the path of an input file.
func (fileNamer *Filenamer) InputFilename() string {
	return filepath.Join(fileNamer.indir, fileNamer.file.Name())
}

// OutputFilename returns the path of an output file.
func (fileNamer *Filenamer) OutputFilename() string {
	return filepath.Join(fileNamer.outdir, fileNamer.file.Name())
}

// MetaFilename returns the path of a meta data file.
func (fileNamer *Filenamer) MetaFilename() string {
	return filepath.Join(fileNamer.outdir, namePart(fileNamer.file.Name())+metaFileSuffix)
}

// Accept accepts files that have the proper extension and that are
// not a directory.
func Accept(file os.FileInfo) bool {
	return !file.IsDir() && hasValidExtension(file)
}

func hasValidExtension(file os.FileInfo) bool {
	return strings.EqualFold(filepath.Ext(file.Name()), extension)
}

func namePart(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

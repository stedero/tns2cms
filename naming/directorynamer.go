package naming

import (
	"os"
)

// DirectoryNamer holds the names of the input- and
// output directories.
type DirectoryNamer struct {
	indir  string
	outdir string
}

// NewDirectoryNamer creates a new instance.
func NewDirectoryNamer(indir string, outdir string) *DirectoryNamer {
	return &DirectoryNamer{indir, outdir}
}

// NewFilenamer creates a Filenamer instance with the directories of this
// DirectoryNamer.
func (directoryNamer *DirectoryNamer) NewFilenamer(file os.FileInfo) *Filenamer {
	return &Filenamer{*directoryNamer, file}
}

// InDir returns the input directory.
func (directoryNamer *DirectoryNamer) InDir() string {
	return directoryNamer.indir
}

// OutDir returns the output directory.
func (directoryNamer *DirectoryNamer) OutDir() string {
	return directoryNamer.outdir
}

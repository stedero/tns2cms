package naming

import (
	"path/filepath"
	"strings"
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

// InDir returns the input directory.
func (directoryNamer *DirectoryNamer) InDir() string {
	return directoryNamer.indir
}

// NewOutdirName creates a new output directory name for the specified path.
func (directoryNamer *DirectoryNamer) NewOutdirName(path string) string {
	source := strings.TrimPrefix(path, directoryNamer.indir)
	dest := filepath.Join(directoryNamer.outdir, source)
	return dest
}

package naming

import (
	"os"
)

type DirectoryNamer struct {
	indir  string
	outdir string
}

func NewDirectoryNamer(indir string, outdir string) *DirectoryNamer {
	return &DirectoryNamer{indir, outdir}
}

func (directoryNamer *DirectoryNamer) NewFilenamer(file os.FileInfo) *Filenamer {
	return &Filenamer{directoryNamer.indir, directoryNamer.outdir, file}
}

func (directoryNamer *DirectoryNamer) InDir() string {
	return directoryNamer.indir
}

func (directoryNamer *DirectoryNamer) OutDir() string {
	return directoryNamer.outdir
}

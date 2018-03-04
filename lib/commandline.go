package lib

import (
	"fmt"
	"os"
)

type DirectoryNamer struct {
	InDir  string
	OutDir string
}

func ParseCommandLine() *DirectoryNamer {
	if len(os.Args) != 3 {
		Usage()
	} else {
		indir, outdir := os.Args[1], os.Args[2]
		if !IsExistingDirectory(indir) {
			fmt.Fprintf(os.Stderr, "%s is not an existing directory\n", indir)
			exit()
		}
		return &DirectoryNamer{indir, outdir}
	}
	return &DirectoryNamer{}
}

func (directoryNamer *DirectoryNamer) NewFilenamer(file os.FileInfo) *Filenamer {
	return &Filenamer{directoryNamer.InDir, directoryNamer.OutDir, file}
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%s <input directory> <output directory>\n", os.Args[0])
	exit()
}

func exit() {
	os.Exit(2)
}

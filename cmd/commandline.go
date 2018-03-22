package cmd

import (
	"fmt"
	"os"
	"ibfd.org/tns2cms/io"
	"ibfd.org/tns2cms/paths"
)

// ParseCommandLine extracts directory names from the command line
func ParseCommandLine() *paths.DirectoryNamer {
	if len(os.Args) != 3 {
		usage()
	} else {
		indir, outdir := os.Args[1], os.Args[2]
		if !io.IsExistingDirectory(indir) {
			fmt.Fprintf(os.Stderr, "%s is not an existing directory\n", indir)
			exit()
		}
		return paths.NewDirectoryNamer(indir, outdir)
	}
	return nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%s <input directory> <output directory>\n", os.Args[0])
	exit()
}

func exit() {
	os.Exit(2)
}

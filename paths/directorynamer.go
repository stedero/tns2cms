package paths

import (
	"path/filepath"
	"strings"
)

var monthsToNum map[string]string

func init() {
	monthsToNum = make(map[string]string)
	monthsToNum["jan"] = "01"
	monthsToNum["feb"] = "02"
	monthsToNum["mar"] = "03"
	monthsToNum["apr"] = "04"
	monthsToNum["may"] = "05"
	monthsToNum["jun"] = "06"
	monthsToNum["jul"] = "07"
	monthsToNum["aug"] = "08"
	monthsToNum["sep"] = "09"
	monthsToNum["oct"] = "10"
	monthsToNum["nov"] = "11"
	monthsToNum["dec"] = "12"
}

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

// InDir returns the input directory name.
func (directoryNamer *DirectoryNamer) InDir() string {
	return directoryNamer.indir
}

// OutDir returns the output directory name.
func (directoryNamer *DirectoryNamer) OutDir() string {
	return directoryNamer.outdir
}

// NewOutdirName creates a new output directory name from the specified input path.
func (directoryNamer *DirectoryNamer) NewOutdirName(inputPath string) string {
	part := subdirName(strings.TrimPrefix(inputPath, directoryNamer.indir))
	return filepath.Join(directoryNamer.outdir, part)
}

func subdirName(part string) string {
	parts := strings.Split(part, "\\")
	month := parts[len(parts)-1]
	if len(month) == 5 {
		lastDir := monthsToNum[strings.ToLower(month[:3])]
		return strings.Join(parts[:len(parts)-1], "\\") + "\\" + lastDir
	}
	return part
}

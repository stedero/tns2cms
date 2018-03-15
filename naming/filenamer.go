package naming

import (
	"os"
	"path/filepath"
	"strings"
)

const extension = ".xml"
const metaFileSuffix = ".metadata.properties" + extension

// Validation defines the possible validation values
type Validation int

// The validate response values
const (
	AcceptDir Validation = iota
	AcceptFile
	RejectDir
	RejectFile
)

// Filenamer holds all data needed to create input- and outputfilenames.
type Filenamer struct {
	DirectoryNamer
	fileName string
}

// NewFilenamerFromRoot creates a filenamer from root directory info.
func NewFilenamerFromRoot(rootDirNamer *DirectoryNamer, currentPath string, fileInfo os.FileInfo) *Filenamer {
	source := strings.TrimSuffix(currentPath, fileInfo.Name())
	part := strings.TrimPrefix(source, rootDirNamer.indir)
	dest := filepath.Join(rootDirNamer.outdir, part)
	return &Filenamer{*NewDirectoryNamer(source, dest), fileInfo.Name()}
}

// NewFilenamer creates a filenamer from root directory info.
func NewFilenamer(dirNamer *DirectoryNamer, fileInfo os.FileInfo) *Filenamer {
	return &Filenamer{*dirNamer, fileInfo.Name()}
}

// Validate determines whether a file or directory can be accepted
func Validate(fileInfo os.FileInfo) Validation {
	if fileInfo.IsDir() {
		if hasValidDirectoryName(fileInfo) {
			return AcceptDir
		}
		return RejectDir
	} else if hasValidExtension(fileInfo) {
		return AcceptFile
	} else {
		return RejectFile
	}
}

// InputFilename returns the path of an input file.
func (fileNamer *Filenamer) InputFilename() string {
	return filepath.Join(fileNamer.indir, fileNamer.fileName)
}

// OutputFilename returns the path of an output file.
func (fileNamer *Filenamer) OutputFilename() string {
	return filepath.Join(fileNamer.outdir, fileNamer.fileName)
}

// MetaFilename returns the path of a meta data file.
func (fileNamer *Filenamer) MetaFilename() string {
	return filepath.Join(fileNamer.outdir, namePart(fileNamer.fileName)+metaFileSuffix)
}

// AcceptDir accepts directories with a valid name.
// func AcceptDir(file os.FileInfo) bool {
// 	return file.IsDir() && hasValidDirectoryName(file)
// }

// AcceptFile accepts files that have the proper extension and that are
// not a directory.
// func AcceptFile(file os.FileInfo) bool {
// 	return !file.IsDir() && hasValidExtension(file)
// }

func hasValidDirectoryName(fileInfo os.FileInfo) bool {
	return !strings.HasPrefix(fileInfo.Name(), ".")
}

func hasValidExtension(fileInfo os.FileInfo) bool {
	return strings.EqualFold(filepath.Ext(fileInfo.Name()), extension)
}

func namePart(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

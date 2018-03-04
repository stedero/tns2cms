package lib

import (
	"os"
	"path/filepath"
	"strings"
)

const extension = ".xml"
const meta_file_suffix = ".metadata.properties" + extension

type Filenamer struct {
	indir  string
	outdir string
	file   os.FileInfo
}

func Accept(file os.FileInfo) bool {
	return filepath.Ext(file.Name()) == extension
}

func (fileNamer *Filenamer) InputFilename() string {
	return filepath.Join(fileNamer.indir, fileNamer.file.Name())
}

func (fileNamer *Filenamer) OutputFilename() string {
	return filepath.Join(fileNamer.outdir, fileNamer.file.Name())
}

func (fileNamer *Filenamer) MetaFilename() string {
	name := strings.TrimSuffix(fileNamer.file.Name(), extension)
	return filepath.Join(fileNamer.outdir, name+meta_file_suffix)
}

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
	return !file.IsDir() && strings.EqualFold(filepath.Ext(file.Name()), extension)
}

func (fileNamer *Filenamer) InputFilename() string {
	return filepath.Join(fileNamer.indir, fileNamer.file.Name())
}

func (fileNamer *Filenamer) OutputFilename() string {
	return filepath.Join(fileNamer.outdir, fileNamer.file.Name())
}

func (fileNamer *Filenamer) MetaFilename() string {
	return filepath.Join(fileNamer.outdir, namePart(fileNamer.file.Name())+meta_file_suffix)
}

func namePart(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

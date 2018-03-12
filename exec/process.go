package exec

import (
	"os"
	"sync"
	"tns2cms/naming"
)

// Process files in chunks of 100 to avoid error "too many open files"
// when processing concurrently
const chunkSize = 100

// Processor defines the process handler.
type Processor struct {
	dirNamer    *naming.DirectoryNamer
	coreHandler func(*naming.Filenamer)
}

// NewProcessor creates a new Processor.
func NewProcessor(dirNamer *naming.DirectoryNamer, coreHandler func(*naming.Filenamer)) *Processor {
	return &Processor{dirNamer, coreHandler}
}

// ExecSequential processes files one by one.
func (processor *Processor) ExecSequential(files []os.FileInfo) {
	for _, file := range files {
		processor.coreHandler(processor.dirNamer.NewFilenamer(file))
	}
}

// ExecConcurrent processes files concurrently
func (processor *Processor) ExecConcurrent(files []os.FileInfo) {
	var waitGroup sync.WaitGroup
	for _, chunk := range chunks(chunkSize, files) {
		waitGroup.Add(len(chunk))
		for _, file := range chunk {
			go processor.exec(file, &waitGroup)
		}
		waitGroup.Wait()
	}
}

func (processor *Processor) exec(file os.FileInfo, waitGroup *sync.WaitGroup) {
	processor.coreHandler(processor.dirNamer.NewFilenamer(file))
	waitGroup.Done()
}

func chunks(chunksize int, files []os.FileInfo) [][]os.FileInfo {
	rows := len(files)/chunksize + 1
	chunks := make([][]os.FileInfo, rows)
	for i := range chunks {
		start := i * chunksize
		end := start + min(chunksize, len(files)-start)
		chunks[i] = files[start:end]
	}
	return chunks
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
